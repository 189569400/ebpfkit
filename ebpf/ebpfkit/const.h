/* SPDX-License-Identifier: GPL-2.0 WITH Linux-syscall-note */
/* Copyright (c) 2020
 *
 * This program is free software; you can redistribute it and/or
 * modify it under the terms of version 2 of the GNU General Public
 * License as published by the Free Software Foundation.
 */
#ifndef _CONST_H_
#define _CONST_H_

#define HTTP_PASS 0
#define HTTP_DROP 1
#define HTTP_EDIT 2

#define HTTP_REQ_PATTERN 61
#define HTTP_REQ_LEN 500
#define HTTP_RESP_LEN 629

#define HTTP_ACTION_HANDLER 0
#define HTTP_ADD_FS_WATCH_HANDLER 1
#define HTTP_DEL_FS_WATCH_HANDLER 2
#define HTTP_GET_FS_WATCH_HANDLER 3
#define DNS_RESP_HANDLER 4

#define DNS_PORT 53
#define DNS_MAX_LENGTH 256
#define DNS_A_RECORD 1
#define DNS_COMPRESSION_FLAG 3

#define LOAD_CONSTANT(param, var) asm("%0 = " param " ll" : "=r"(var))

__attribute__((always_inline)) static u16 load_http_server_port() {
    u64 http_server_port = 0;
    LOAD_CONSTANT("http_server_port", http_server_port);
    return (u16)http_server_port;
}

#endif
