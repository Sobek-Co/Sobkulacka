//======================================================================//
// Purpose:     Math Library
//
// Authors:     David Sobek <xsobekd00@fit.vutbr.cz>
// Date:        2026-03-15
//======================================================================//
/**
 * @file ourmath.c
 * @author David Sobek <xsobekd00@fit.vutbr.cz>
 * 
 * @brief Math Library.
 */

#include <stdlib.h>
#include "ourmath.h"

m_state* m_init_state() {
    m_state* state = (m_state*)malloc(sizeof(m_state));
    state->current = 0.0;
    state->error = SUCESS;
    return state;
}

void m_clear(m_state* state) {}
double m_get_state(const m_state* state) {}
m_err m_get_err(const m_state* state) {}
void m_set_state(m_state* state, double value) {}
void m_add(m_state* state, double value) {}
void m_sub(m_state* state, double value) {}
void m_mul(m_state* state, double value) {}
void m_div(m_state* state, double value) {}
void m_mod(m_state* state, int value) {}
void m_power(m_state* state, int value) {}
void m_root(m_state* state, int value) {}
void m_factorial(m_state* state) {}
void m_sin(m_state* state) {}
void m_cos(m_state* state) {}
void m_tg(m_state* state) {}
void m_cotg(m_state* state) {}