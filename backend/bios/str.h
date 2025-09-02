#include <stdlib.h>
#include <stdio.h>

typedef struct {
    char *data;
    size_t len;
    size_t cap;
} String;

String* str_new();
void str_append(String *s, char c);
void str_init(String *s);
