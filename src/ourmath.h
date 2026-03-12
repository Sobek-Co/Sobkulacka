//======================================================================//
// Purpose:     Math Library
//
// Authors:     Vojtěch Vítek <xvitekv00@fit.vutbr.cz>
// Date:        2026-03-15
//======================================================================//
/**
 * @file ourmath.h
 * @author Vojtěch Vítek <xvitekv00@fit.vutbr.cz>
 * 
 * @brief Header file for the Math Library.
 */

#ifndef OURMATH_H
#define OURMATH_H

typedef enum {
    SUCESS,
    DIVISION_BY_ZERO,
} err;

typedef struct {
    double current;
    err error;
} state;

void m_init_state(state* state);
void m_clear(state* state);
double m_get_state(const state* state);

void m_add(state* state, double value);         // + sobek
void m_sub(state* state, double value);         // - sobek
void m_mul(state* state, double value);         // * sobek
void m_div(state* state, double value);         // / sobek
void m_mod(state* state, int value);            // %
void m_power(state* state, int value);          // **
void m_root(state* state, int value);           // odmocnina
void m_factorial(state* state, double value);   // problem budouciho kubizona

void m_sin(state* state, double value);
void m_cos(state* state, double value);
void m_tg(state* state, double value);
void m_cotg(state* state, double value);

#endif