#include "heif.h"

void print_hex(void *buf, int len)
{
    for (int i = 0; i < len; i++) {
        if (i && !(i % 16))
        {
            printf("\n");
        }
        else if (i && !(i % 4)) printf("   ");
        printf("%02X ", ((unsigned char *)buf)[i]);
    }
    printf("\n");
}

void print_str(void *buf, int len)
{
    for (int i = 0; i < len; i++)
    {
        printf("%c", ((char *)buf)[i]);
    }
}

int min(int foo, int bar)
{
    return foo < bar ? foo : bar;
}

int test(const char *src)
{
    struct heif_t box_head;
    
    struct ftyp_t ftyp_box;
    int fd = open(src, O_RDONLY);
    pread(fd, &ftyp_box, sizeof(struct ftyp_t), 0);

    print_str(&(ftyp_box.head.type), 4);
    printf(" ");
    print_str(&(ftyp_box.major_brand), 4);

    uint32_t siz = ntohl(ftyp_box.head.size);
    uint32_t ver = ntohl(ftyp_box.minor_version);
    printf("\nlen=%d, version=%d\n", siz, ver);

    off_t off = 0;
    for (int i = 0; i < 4; i++)
    {
        if (siz < 8)
        {
            siz = 8;
        }
        off += siz;
        pread(fd, &box_head, sizeof(struct heif_t), off);
        print_hex(&(box_head.size), 4);
        printf(" ");
        siz = ntohl(box_head.size);
        print_str(&box_head.type, 4);
        printf(" len=%d\n", siz);
        if (*((int32_t *)"mdat") == box_head.type)
        {
            char buf[64];
            pread(fd, buf, 64, off + 8);
            print_hex(buf, 64);
            print_str(buf, 64);
        }
    }

    // for (int i = 0; i < len; i += 4)
    // {
    //     print_str(buf + i, min(len - i, 4));
    //     printf(" ");
    // }
    // printf("\n");
    // print_hex(buf, len);


    close(fd);
    return 0;
}