//======================================================================//
// Purpose:     Math Library header
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

/*
 * @brief Error types for the Math Library.
 */
typedef enum {
    SUCCESS,
    DIVISION_BY_ZERO,
    NON_INTEGER_NUMBER,
    NON_POSITIVE_NUMBER,
    NEGATIVE_ROOT,
} m_err;

/*
 * @brief State structure for the Math Library.
 */
typedef struct {
    double current;
    double accuracy;
    m_err error;
} m_state;

/*
 * @brief Initializes the state structure.
 *
 * @return A pointer to the initialized state structure.
*/
m_state* m_init_state();

/*
 * @brief Clears the state structure, resetting the current value to 0 and error status to SUCESS.
 *
 * @param state Pointer to the state structure to be cleared.
*/
void m_clear(m_state* state);

/*
 * @brief Frees the memory allocated for the state structure.
 *
 * @param state Pointer to the state structure to be freed.
 */
void m_free_state(m_state* state);

/*
 * @brief Gets the current value from the state structure.
 *
 * @param state Pointer to the state structure.
 * @return The current value.
 */
double m_get_state(const m_state* state);

/*
 * @brief Gets the accuracy from the state structure.
 *
 * @param state Pointer to the state structure.
 * @return The accuracy value (double).
 */
double m_get_accuracy(const m_state* state);

/*
 * @brief Gets the current error status from the state structure.
 *
 * @param state Pointer to the state structure.
 * @return The current error status.
 */
m_err m_get_err(const m_state* state);

/*
 * @brief Sets the current value in the state structure.
 *
 * @param state Pointer to the state structure.
 * @param value The value to be set.
 */
void m_set_state(m_state* state, double value);

/*
 * @brief Sets the accuracy in the state structure.
 *
 * @param state Pointer to the state structure.
 * @param accuracy The accuracy value to be set.
 */
void m_set_accuracy(m_state* state, double accuracy);

/*
 * @brief Calculates the arctangent of the current value in the state structure.
 *
 * @param state Pointer to the state structure.
 */
void m_arctan(m_state* state);

/*
 * @brief Sets the current value to π, acoording to accuracy defined in the state structure.
 *
 * @param state Pointer to the state structure.
 */
void m_pi(m_state* state);

/*
 * @brief Rounds the current value in the state structure.
 *
 * @param state Pointer to the state structure.
 */
void m_round(m_state* state);

/*
 * @brief Adds a value to the current value in the state structure.
 *
 * @param state Pointer to the state structure.
 * @param value The value to be added.
 */
void m_add(m_state* state, double value);

/*
 * @brief Subtracts a value from the current value in the state structure.
 *
 * @param state Pointer to the state structure.
 * @param value The value to be subtracted.
 */
void m_sub(m_state* state, double value);

/*
 * @brief Multiplies the current value in the state structure by a value.
 *
 * @param state Pointer to the state structure.
 * @param value The value to multiply by.
 */
void m_mul(m_state* state, double value);

/*
 * @brief Divides the current value in the state structure by a value.
 *
 * @param state Pointer to the state structure.
 * @param value The value to divide by.
 */
void m_div(m_state* state, double value);

/*
 * @brief Calculates the remainder of the current value in the state structure when divided by a value.
 *
 * @param state Pointer to the state structure.
 * @param value The value to divide by.
 */
void m_mod(m_state* state, int value);

/*
 * @brief Raises the current value in the state structure to a power.
 *
 * @param state Pointer to the state structure.
 * @param value The power to raise to.
 */
void m_power(m_state* state, int value);

/*
 * @brief Calculates the root of the current value in the state structure.
 *
 * @param state Pointer to the state structure.
 * @param value The root to calculate.
 */
void m_root(m_state* state, int value);

/*
 * @brief Calculates the factorial of the current value in the state structure.
 *
 * @param state Pointer to the state structure.
 */
void m_factorial(m_state* state);

/*
 * @brief Calculates the sine of the current value in the state structure.
 *
 * @param state Pointer to the state structure. The value is interpreted in degrees.
 */
void m_sin(m_state* state);

/*
 * @brief Calculates the cosine of the current value in the state structure.
 *
 * @param state Pointer to the state structure. The value is interpreted in degrees.
 */
void m_cos(m_state* state);

/*
 * @brief Calculates the tangent of the current value in the state structure.
 *
 * @param state Pointer to the state structure. The value is interpreted in degrees.
 */
void m_tg(m_state* state);

/*
 * @brief Calculates the cotangent of the current value in the state structure.
 *
 * @param state Pointer to the state structure. The value is interpreted in degrees.
 */
void m_cotg(m_state* state);

#endif