#ifndef _HEIF_H_
#define _HEIF_H_

#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <unistd.h>
#include <stdint.h>
#include <fcntl.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <sys/sendfile.h>
#include <arpa/inet.h>

struct heif_t
{
    uint32_t size;
    uint32_t type;
};

struct ftyp_t
{
    struct heif_t head;
    uint32_t major_brand;
    uint32_t minor_version;
};

void print_hex(void *buf, int len);

void print_str(void *buf, int len);

int min(int foo, int bar);

int test(const char *src);

#endif
