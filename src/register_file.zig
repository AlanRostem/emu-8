const t = @import("typedef.zig");
const RegisterFile = struct {
    general_purpose_registers: t.Byte[0x10],
    index_register: t.Word,
    program_counter: t.Word,
    stack_pointer: t.Word,
    delay_timer_register: t.Byte,
    sound_timer_register: t.Byte,
    current_operand: t.Word,
};
