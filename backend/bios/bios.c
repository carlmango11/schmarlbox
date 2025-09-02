#include <stdint.h>
#include <string.h>
#include "str.h"

const char KEY_CR = 0xd;
const int DISPLAY_ADDR = 0x2200;
const int INPUT_ADDR = 0x3000;

void write_char(const char c) {
    volatile uint8_t* addr = (volatile uint8_t*)DISPLAY_ADDR;
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

String* listen_command(void) {
    String *command = str_new();

    volatile uint8_t* input = (volatile uint8_t*)INPUT_ADDR;
    int i = 0;

    print("\r\n> ");

    while (1) {
        const uint8_t val = *input;

        if (val == 0) {
            continue;
        }

        if (val == 'Q') {
            return NULL;
        }

        if (val == KEY_CR) {
            return command;
        }

        process_char(val);
        str_append(command, (char)val);
    }
}

void main(void) {
    String *command;
    char welcome[] = "Welcome to SchmarlBox";

    print("\033[2J\033[H\033[1;32m");
    print(welcome);

    while (1) {
        command = listen_command();
        if (command == NULL) {
            break;
        }

        if (strcmp(command->data, "quit") == 0) {
            exit(0);
        }
        exit(1);

        print("\r\nExecuting: ");
        print(command->data);
    }
}