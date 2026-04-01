//======================================================================//
// Purpose:     Tests for Math Library
//
// Authors:     David Sobek <xsobekd00@fit.vutbr.cz>
// Date:        2026-03-13
//======================================================================//
/**
 * @file test_ourmath.c
 * @author David Sobek <xsobekd00@fit.vutbr.cz>
 * 
 * @brief Tests for the Math Library.
 */

#include <check.h>
#include <math.h>
#include "ourmath.h"

// ---------------------------
// Helper for comparing doubles
// ---------------------------
int double_equal(double a, double b) {
    return fabs(a - b) < 1e-9;
}

// ---------------------------
// Helper for comparing doubles with less accuracy
// ---------------------------
int double_equal_forgiving(double a, double b) {
    return fabs(a - b) < 1e-3;
}

// ---------------------------
// m_init_state
// ---------------------------
START_TEST(test_m_init_state) {
    m_state* s = m_init_state();
    ck_assert(s != NULL);
    ck_assert(double_equal(m_get_state(s), 0.0));
    ck_assert(m_get_err(s) == SUCCESS);
    m_free_state(s);
}
END_TEST

// ---------------------------
// m_clear
// ---------------------------
START_TEST(test_m_clear) {
    m_state* s = m_init_state();
    m_set_state(s, 5.0);
    m_clear(s);
    ck_assert(double_equal(m_get_state(s), 0.0));
    ck_assert(m_get_err(s) == SUCCESS);
    m_free_state(s);
}
END_TEST

// ---------------------------
// m_set_state / m_get_state
// ---------------------------
START_TEST(test_m_set_get_state) {
    m_state* s = m_init_state();
    m_set_state(s, 3.14159);
    ck_assert(double_equal(m_get_state(s), 3.14159));
    m_free_state(s);
}
END_TEST

// ---------------------------
// m_set_accuracy / m_get_accuracy
// ---------------------------
START_TEST(test_m_set_get_accuracy) {
    m_state* s = m_init_state();
    m_set_accuracy(s, 1e-9);
    ck_assert(double_equal(m_get_accuracy(s), 1e-9));
    m_free_state(s);
}
END_TEST

// ---------------------------
// m_get_err
// ---------------------------
START_TEST(test_m_get_err) {
    m_state* s = m_init_state();
    ck_assert(m_get_err(s) == SUCCESS);
    m_div(s, 0.0);
    ck_assert(m_get_err(s) == DIVISION_BY_ZERO);
    m_free_state(s);
}
END_TEST

// ---------------------------
// m_add
// ---------------------------
START_TEST(test_m_add) {
    m_state* s = m_init_state();
    m_add(s, 2.5);
    ck_assert(double_equal(m_get_state(s), 2.5));
    m_free_state(s);
}
END_TEST

// ---------------------------
// m_sub
// ---------------------------
START_TEST(test_m_sub) {
    m_state* s = m_init_state();
    m_set_state(s, 5.0);
    m_sub(s, 3.0);
    ck_assert(double_equal(m_get_state(s), 2.0));
    m_free_state(s);
}
END_TEST

// ---------------------------
// m_mul
// ---------------------------
START_TEST(test_m_mul) {
    m_state* s = m_init_state();
    m_set_state(s, 3.0);
    m_mul(s, 4.0);
    ck_assert(double_equal(m_get_state(s), 12.0));
    m_free_state(s);
}
END_TEST

// ---------------------------
// m_div
// ---------------------------
START_TEST(test_m_div) {
    m_state* s = m_init_state();
    m_set_state(s, 10.0);
    m_div(s, 2.0);
    ck_assert(double_equal(m_get_state(s), 5.0));

    m_div(s, 0.0);
    ck_assert(m_get_err(s) == DIVISION_BY_ZERO);
    m_free_state(s);
}
END_TEST

// ---------------------------
// m_mod
// ---------------------------
START_TEST(test_m_mod) {
    m_state* s = m_init_state();
    m_set_state(s, 10.0);
    m_mod(s, 3);
    ck_assert(double_equal(m_get_state(s), 1.0));
    m_free_state(s);
}
END_TEST

// ---------------------------
// m_power
// ---------------------------
START_TEST(test_m_power) {
    m_state* s = m_init_state();
    m_set_state(s, 2.0);
    m_power(s, 3);
    ck_assert(double_equal(m_get_state(s), 8.0));
    m_free_state(s);
}
END_TEST

// ---------------------------
// m_root
// ---------------------------
START_TEST(test_m_root) {
    m_state* s = m_init_state();
    m_set_state(s, 9.0);
    m_root(s, 2);
    ck_assert(double_equal(m_get_state(s), 3.0));
    m_free_state(s);
}
END_TEST

// ---------------------------
// m_factorial
// ---------------------------
START_TEST(test_m_factorial) {
    m_state* s = m_init_state();
    m_set_state(s, 5.0);
    m_factorial(s);
    ck_assert(double_equal(m_get_state(s), 120.0));
    m_free_state(s);
}
END_TEST

// ---------------------------
// m_sin
// ---------------------------
START_TEST(test_m_sin) {
    m_state* s = m_init_state();
    m_set_state(s, 90.0);
    m_sin(s);
    ck_assert(double_equal(m_get_state(s), 1.0));
    m_free_state(s);
}
END_TEST

// ---------------------------
// m_cos
// ---------------------------
START_TEST(test_m_cos) {
    m_state* s = m_init_state();
    m_set_state(s, 180.0);
    m_cos(s);
    ck_assert(double_equal(m_get_state(s), -1.0));
    m_free_state(s);
}
END_TEST

// ---------------------------
// m_tg
// ---------------------------
START_TEST(test_m_tg) {
    m_state* s = m_init_state();
    m_set_state(s, 45.0);
    m_tg(s);
    ck_assert(double_equal(m_get_state(s), 1.0));
    m_free_state(s);
}
END_TEST

// ---------------------------
// m_cotg
// ---------------------------
START_TEST(test_m_cotg) {
    m_state* s = m_init_state();
    m_set_state(s, 45.0);
    m_cotg(s);
    ck_assert(double_equal(m_get_state(s), 1.0));
    m_free_state(s);
}
END_TEST

// ---------------------------
// Suite setup
// ---------------------------
Suite* ourmath_suite(void) {
    Suite* s = suite_create("OurMath");

    TCase* tc_init = tcase_create("Initialization");
    tcase_add_test(tc_init, test_m_init_state);
    tcase_add_test(tc_init, test_m_clear);
    tcase_add_test(tc_init, test_m_set_get_state);
    tcase_add_test(tc_init, test_m_set_get_accuracy);
    tcase_add_test(tc_init, test_m_get_err);
    suite_add_tcase(s, tc_init);

    TCase* tc_arith = tcase_create("Arithmetic");
    tcase_add_test(tc_arith, test_m_add);
    tcase_add_test(tc_arith, test_m_sub);
    tcase_add_test(tc_arith, test_m_mul);
    tcase_add_test(tc_arith, test_m_div);
    tcase_add_test(tc_arith, test_m_mod);
    tcase_add_test(tc_arith, test_m_power);
    tcase_add_test(tc_arith, test_m_root);
    tcase_add_test(tc_arith, test_m_factorial);
    suite_add_tcase(s, tc_arith);

    TCase* tc_trig = tcase_create("Trig");
    tcase_add_test(tc_trig, test_m_sin);
    tcase_add_test(tc_trig, test_m_cos);
    tcase_add_test(tc_trig, test_m_tg);
    tcase_add_test(tc_trig, test_m_cotg);
    suite_add_tcase(s, tc_trig);

    return s;
}

// ---------------------------
// Main runner
// ---------------------------
int main(void) {
    Suite* s = ourmath_suite();
    SRunner* sr = srunner_create(s);

    srunner_run_all(sr, CK_NORMAL);
    int nf = srunner_ntests_failed(sr);
    srunner_free(sr);

    return nf == 0 ? 0 : 1;
}