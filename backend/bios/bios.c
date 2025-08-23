#include <stdint.h>

void write_char(const char c) {
    volatile uint8_t* addr = (volatile uint8_t*)0x2000;
    *addr = (volatile uint8_t) c;
}

void print(const char *str) {
    int i = 0;
    while (str[i] != '\0') {
        write_char(str[i]);
        i++;
    }
}

void process_char(const uint8_t c) {
    write_char(c);
}

void main(void) {
    volatile uint8_t* input = (volatile uint8_t*)0x3000;

    char welcome[] = "\nWelcome to SchmarlBox";
    print(welcome);

    while (1) {
        const uint8_t val = *input;

        if (val == 0) {
            continue;
        }

        process_char(val);
    }
}