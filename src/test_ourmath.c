// test_ourmath.c
#include <check.h>
#include "ourmath.h"

// Example test for a function add(a, b)
START_TEST(test_add)
{
    m_state state;
    m_clear(&state);
    ck_assert_int_eq(m_get_state(&state), 0);
    ck_assert_int_eq(m_get_err(&state), SUCESS);
}
END_TEST

// Example test for subtract
START_TEST(test_subtract)
{
    ck_assert_int_eq(5-3, 2);
    ck_assert_int_eq(0-4, -4);
}
END_TEST

// Suite definition
Suite* ourmath_suite(void)
{
    Suite *s;
    TCase *tc_core;

    s = suite_create("OurMath");

    tc_core = tcase_create("Core");

    tcase_add_test(tc_core, test_add);
    tcase_add_test(tc_core, test_subtract);

    suite_add_tcase(s, tc_core);

    return s;
}

int main(void)
{
    int failed;
    Suite *s;
    SRunner *sr;

    s = ourmath_suite();
    sr = srunner_create(s);

    srunner_run_all(sr, CK_NORMAL);
    failed = srunner_ntests_failed(sr);
    srunner_free(sr);

    return (failed == 0) ? 0 : 1;
}