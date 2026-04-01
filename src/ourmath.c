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
    state->error = SUCCESS;
    return state;
}

void m_clear(m_state* state) {
    state->current = 0.0;
    state->error = SUCCESS;
}

void m_free_state(m_state* state) {
    if (state != NULL) {
        free(state);
    }
}

double m_get_state(const m_state* state) {
    return state->current;
}

double m_get_accuracy(const m_state* state) {
    return state->accuracy;
}

m_err m_get_err(const m_state* state) {
    return state->error;
}

void m_set_state(m_state* state, double value) {
    state->current = value;
    state->error = SUCCESS;
}

void m_pi(m_state* state) {
    double pi = 0.0;
    double term;
    int n = 0;

    do {
        term = (1.0 / (n * 4.0 + 1.0)) - (1.0 / (n * 4.0 + 3.0));
        pi += term;
        n++;
    } while (term > state->accuracy || term < -state->accuracy);

    state->current = pi * 4.0;
    state->error = SUCCESS;
}

void m_round(m_state* state) {
    // doplnit
}

void m_add(m_state* state, double value) {
    state->current += value;
    state->error = SUCCESS;
}

void m_sub(m_state* state, double value) {
    state->current -= value;
    state->error = SUCCESS;
}
void m_mul(m_state* state, double value) {
    state->current = state->current * value;
    state->error = SUCCESS;
}

void m_div(m_state* state, double value) {
    if (value == 0) {
        state->error = DIVISION_BY_ZERO;
        return;
    }

    if (value == 1) {
        state->error = SUCCESS;
        return;
    }

    state->current = state->current / value;
    state->error = SUCCESS;
}

void m_mod(m_state* state, int value) {
    if (state->current != (int)state->current) {
        state->error = NON_INTEGER_NUMBER;
        return;
    }

    if (value == 1) {
        state->error = SUCCESS;
        state->current = 0;
        return;
    }

    state->current = (double)((int)state->current % value);
    state->error = SUCCESS;
}

void m_power(m_state* state, int value) {
    if (value == 0) {
        state->current = 1;
        state->error = SUCCESS;
        return;
    }

    if (value == 0) {
        state->error = SUCCESS;
        return;
    }

    if (value > 0) {
        double base = state->current;
        for (int i = value - 1; i > 0; i--) {
            m_mul(state, base);
        }
        state->error = SUCCESS;
        return;
    }

    if (value < 0) {
        m_state* temp_state = m_init_state();
        m_power(temp_state, -value);

        state->current = 1 / temp_state->current;
        state->error = SUCCESS;

        m_free_state(temp_state);
        return;
    }
}

void m_root(m_state* state, int value) {
    if (value <= 0) {
        state->error = NON_POSITIVE_NUMBER;
        return;
    }

    if (state->current < 0 && value % 2 == 0) {
        state->error = NEGATIVE_ROOT;
        return;
    }

    double aproximate = state->current > 1 ? state->current : 1;
    double prev;
    double eps = state->accuracy;

    m_state* temp = m_init_state();
    do {
        prev = aproximate;

        m_set_state(temp, aproximate);
        m_power(temp, value - 1);

        aproximate = ((value - 1) * aproximate + state->current / temp->current) / value;
    } while ((aproximate - prev > eps) || (prev - aproximate > eps));
    m_free_state(temp);

    state->current = aproximate;
    state->error = SUCCESS;
}

void m_factorial(m_state* state) {
    if (state->current != (int)state->current) {
        state->error = NON_INTEGER_NUMBER;
        return;
    }

    if (state->current < 0) {
        state->error = NON_POSITIVE_NUMBER;
        return;
    }

    if (state->current == 0) {
        state->current = 1;
        state->error = SUCCESS;
    }

    for (int multiplier = state->current - 1; multiplier > 1; multiplier--) {
        m_mul(state, multiplier);
    }
    state->error = SUCCESS;
}

void m_sin(m_state* state) {
    m_state* temp_state = m_init_state();
    m_pi(temp_state);
    double radians = state->current * temp_state->current / 180.0;
    m_free_state(temp_state);

    double term = radians;
    double sum = term;
    int n = 1;
    
    while (term > state->accuracy|| term < -state->accuracy) {
        term *= -radians * radians / ((n + 1) * (n + 2));
        sum += term;
        n += 2;
    }
    
    state->current = sum;
    state->error = SUCCESS;
}

void m_cos(m_state* state) {
    m_state* temp_state = m_init_state();
    m_pi(temp_state);
    double radians = state->current * temp_state->current / 180.0;
    m_free_state(temp_state);

    double term = 1.0;
    double sum = term;
    int n = 0;

    while (term > state->accuracy || term < -state->accuracy) {
        term *= -radians * radians / ((n + 1) * (n + 2));
        sum += term;
        n += 2;
    }

    state->current = sum;
    state->error = SUCCESS;
}

void m_tg(m_state* state) {
    m_state* temp_state = m_init_state();

    m_sin(temp_state);
    double sin_value = temp_state->current;
    m_cos(temp_state);
    double cos_value = temp_state->current;
    m_free_state(temp_state);

    if (cos_value == 0) {
        state->error = DIVISION_BY_ZERO;
        return;
    }

    state->current = sin_value / cos_value;
    state->error = SUCCESS;
}

void m_cotg(m_state* state) {
    m_state* temp_state = m_init_state();
    
    m_sin(temp_state);
    double sin_value = temp_state->current;
    m_cos(temp_state);
    double cos_value = temp_state->current;
    m_free_state(temp_state);

    if (sin_value == 0) {
        state->error = DIVISION_BY_ZERO;
        return;
    }

    state->current = cos_value / sin_value;
    state->error = SUCCESS;
}