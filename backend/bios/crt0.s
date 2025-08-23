    .export _init
    .export __STARTUP__ : absolute = 1 ; mark this as the startup code
    .import _main

.segment "STARTUP"

_init:
    ldx #$FF ; stack pointer
    txs

    cld

    jsr _main
    brk

catch:
    jmp catch

irq_handler:
    rti

nmi_handler:
    rti

; TODO: split to separate file
.segment "VECTORS"
    .word nmi_handler
    .word _init
    .word irq_handler