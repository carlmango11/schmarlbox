    .export _init
    .export _exit
    .export __STARTUP__ : absolute = 1 ; mark this as the startup code
    .import _main
.import    copydata, zerobss, initlib, donelib

.include  "zeropage.inc"

.segment "STARTUP"

_exit:
    brk

_init:
    ldx #$FF ; stack pointer
    txs
          JSR     zerobss              ; Clear BSS segment
          JSR     copydata             ; Initialize DATA segment
          JSR     initlib              ; Run constructors

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