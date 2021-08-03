use rdp::core::client::Connector;
use rdp::core::event::RdpEvent;
use rdp::model::error::*;
use std::net::TcpStream;

#[repr(C)]
pub struct CGOString {
    data: *mut u8,
    len: u16,
}

impl From<CGOString> for String {
    fn from(s: CGOString) -> String {
        unsafe { String::from_raw_parts(s.data, s.len.into(), s.len.into()) }
    }
}

#[no_mangle]
pub extern "C" fn connect_rdp(
    go_addr: CGOString,
    go_username: CGOString,
    go_password: CGOString,
    screen_width: u16,
    screen_height: u16,
) {
    println!("RDP client start");

    // Convert from C to Rust types.
    let addr = String::from(go_addr);
    let username = String::from(go_username);
    let password = String::from(go_password);
    println!("parsed addr {}", addr);

    // Connect and authenticate.
    let tcp = TcpStream::connect(addr).unwrap();
    println!("connected TCP");
    let mut connector = Connector::new()
        .screen(screen_width, screen_height)
        .credentials(".".to_string(), username.to_string(), password.to_string());
    let mut client = connector.connect(tcp).unwrap();

    // Read incoming events.
    loop {
        if let Err(Error::RdpError(e)) = client.read(|rdp_event| match rdp_event {
            RdpEvent::Bitmap(bitmap) => {
                println!("got bitmap {}x{}", bitmap.width, bitmap.height);
            }
            RdpEvent::Pointer(pointer) => {
                println!("got pointer x: {} y: {}", pointer.x, pointer.y);
            }
            RdpEvent::Key(key) => {
                println!("got key code {}", key.code);
            }
        }) {
            match e.kind() {
                RdpErrorKind::Disconnect => {
                    println!("server ask for disconnect");
                }
                _ => println!("{:?}", e),
            }
            break;
        }
    }

    client.shutdown().unwrap();

    println!("RDP client done");
}
