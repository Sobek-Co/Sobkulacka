// ourmath.c
#include <stdlib.h>
#include "ourmath.h"

m_state* m_init_state() {
    m_state* state = (m_state*)malloc(sizeof(m_state));
    if (state) {
        state->current = 0.0;
        state->error = SUCESS;
    }
    return state;
}

void m_clear(m_state* state) {
    if (state) {
        state->current = 0.0;
        state->error = SUCESS;
    }
}

double m_get_state(const m_state* state) {
    return state ? state->current : 0.0;
}

m_err m_get_err(const m_state* state) {
    return state ? state->error : SUCESS;
}

void m_set_state(m_state* state, double value) {
    if (state) state->current = value;
}

void m_add(m_state* state, double value) { (void)state; (void)value; }
void m_sub(m_state* state, double value) { (void)state; (void)value; }
void m_mul(m_state* state, double value) { (void)state; (void)value; }
void m_div(m_state* state, double value) { (void)state; (void)value; }
void m_mod(m_state* state, int value) { (void)state; (void)value; }
void m_power(m_state* state, int value) { (void)state; (void)value; }
void m_root(m_state* state, int value) { (void)state; (void)value; }
void m_factorial(m_state* state) { (void)state; }
void m_sin(m_state* state) { (void)state; }
void m_cos(m_state* state) { (void)state; }
void m_tg(m_state* state) { (void)state; }
void m_cotg(m_state* state) { (void)state; }