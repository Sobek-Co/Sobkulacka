#include <stdio.h>
#include "../ourmath.h"

int main(void) {
    double x;
    int n = 0;

    // Initialize all state structures
    m_state *sum = m_init_state();
    m_state *sum_sq = m_init_state();
    m_state *tmp = m_init_state();
    m_state *result = m_init_state();

    if (!sum || !sum_sq || !tmp || !result) {
        fprintf(stderr, "Memory allocation failed\n");

        if (sum) m_free_state(sum);
        if (sum_sq) m_free_state(sum_sq);
        if (tmp) m_free_state(tmp);
        if (result) m_free_state(result);

        return 1;
    }

    // Set accuracy
    m_set_accuracy(sum, 1e-9);
    m_set_accuracy(sum_sq, 1e-9);
    m_set_accuracy(tmp, 1e-9);
    m_set_accuracy(result, 1e-9);

    // Initialize values
    m_set_state(sum, 0.0);
    m_set_state(sum_sq, 0.0);
    m_set_state(result, 0.0);

    // Read input until EOF
    while (scanf("%lf", &x) == 1) {
        n++;

        // sum += x
        m_add(sum, x);

        // sum_sq += x^2
        m_set_state(tmp, x);
        m_power(tmp, 2);
        m_add(sum_sq, m_get_state(tmp));
    }

    if (n < 2) {
        fprintf(stderr, "Not enough data\n");

        m_free_state(sum);
        m_free_state(sum_sq);
        m_free_state(tmp);
        m_free_state(result);

        return 1;
    }

    // result = sum_sq
    m_set_state(result, m_get_state(sum_sq));

    // tmp = (sum^2) / n
    m_set_state(tmp, m_get_state(sum));
    m_power(tmp, 2);
    m_div(tmp, (double)n);

    // result = sum_sq - (sum^2 / n)
    m_sub(result, m_get_state(tmp));

    // result = result / (n - 1)
    m_div(result, (double)(n - 1));

    // sqrt(result)
    m_root(result, 2);

    // Output result
    printf("%.10f\n", m_get_state(result));

    // Free memory
    m_free_state(sum);
    m_free_state(sum_sq);
    m_free_state(tmp);
    m_free_state(result);

    return 0;
}