#include <stdlib.h>
#include <stdio.h>
#include <stdint.h>
#include <string.h>
#include "str.h"

const size_t INIT_SIZE = 16;

String* str_new() {
    String *s = malloc(sizeof(String));

    s->len = 0;
    s->cap = INIT_SIZE;

    s->data = malloc(s->cap);
    if (s->data == NULL) {
        exit(1);
    }

    s->data[0] = '\0';

    return s;
}

void str_append(String *s, char c) {
    if (s->len == s->cap) {
        s->cap *= s->cap;
        s->data = realloc(s->data, s->cap);

        if(s->data == NULL) {
            exit(1);
        }
    }

    s->data[s->len] = c;
    s->len++;
    s->data[s->len] = '\0';
}

void str_free(String *s) {
    free(s->data);
    s->data = NULL;
    s->cap = 0;
    s->len = 0;
}
