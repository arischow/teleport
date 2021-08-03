#include <stdarg.h>
#include <stdbool.h>
#include <stdint.h>
#include <stdlib.h>

typedef struct CGOString {
  uint8_t *data;
  uint16_t len;
} CGOString;

void connect_rdp(struct CGOString go_addr,
                 struct CGOString go_username,
                 struct CGOString go_password,
                 uint16_t screen_width,
                 uint16_t screen_height);
