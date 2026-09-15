#pragma once

void sys_exit() {
    __asm__ __volatile__ ("ebreak");
}
