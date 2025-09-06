.export _init
.export _exit
.export _enable_paging
.export _disable_paging
.export __STARTUP__ : absolute = 1 ; mark this as the startup code
.import _main
.import    copydata, zerobss, initlib, donelib

.include  "zeropage.inc"

.segment "STARTUP"

.proc _enable_paging
    .byte $02
    rts
.endproc

.proc _disable_paging
    .byte $03
    rts
.endproc

_exit:
    brk

_init:
    ldx #$FF ; stack pointer
    txs

    jsr     zerobss              ; Clear BSS segment
    jsr     copydata             ; Initialize DATA segment
    jsr     initlib              ; Run constructors

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