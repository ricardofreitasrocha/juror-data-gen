--DELETE FROM juror_mod.accused;
DELETE FROM juror_mod.payment_data;
DELETE FROM juror_mod.app_setting;
DELETE FROM juror_mod.appearance;
DELETE FROM juror_mod.appearance_audit;
DELETE FROM juror_mod.financial_audit_details_appearances;
DELETE FROM juror_mod.financial_audit_details;
DELETE FROM juror_mod.bulk_print_data;
DELETE FROM juror_mod.contact_log;
DELETE FROM juror_mod.coroner_pool_detail;
DELETE FROM juror_mod.coroner_pool;
DELETE FROM juror_mod.court_location_audit;
DELETE FROM juror_mod.court_catchment_area;
--DELETE FROM juror_mod.court_region;
DELETE FROM juror_mod.expense_rates;
DELETE FROM juror_mod.juror_audit;
DELETE FROM juror_mod.juror_history;
DELETE FROM juror_mod.pool_comments;
DELETE FROM juror_mod.pool_history;
DELETE FROM juror_mod.juror_pool;
DELETE FROM juror_mod.pool;
DELETE FROM juror_mod.message;

DELETE FROM juror_mod.user_juror_response_audit;
DELETE FROM juror_mod.juror_response_aud;
DELETE FROM juror_mod.juror_response_cjs_employment;
DELETE FROM juror_mod.juror_reasonable_adjustment;
DELETE FROM juror_mod.juror_response;
--
DELETE FROM juror_mod.juror;

delete from juror_mod.voters;

--DELETE FROM juror_mod.message;
--DELETE FROM juror_mod.notify_template_field;
--DELETE FROM juror_mod.notify_template_mapping;
--DELETE FROM juror_mod.password;
--DELETE FROM juror_mod.payment_data;

--DELETE FROM juror_mod.region_notify_template;

--DELETE FROM juror_mod.user_roles;
--DELETE FROM juror_mod.user_courts;
--DELETE FROM juror_mod.users;

delete from juror_mod.users_audit;
DELETE FROM juror_mod.rev_info;
--DELETE FROM JUROR_DIGITAL.STAFF_AUDIT;

--DELETE FROM juror_mod.utilisation_stats;

DELETE FROM juror_mod.juror_trial;
DELETE FROM juror_mod.trial;
DELETE FROM juror_mod.judge;

UPDATE juror_mod.court_location set assembly_room = null;
DELETE FROM juror_mod.courtroom;
DELETE FROM juror_mod.pending_juror;
DELETE FROM juror_mod.welsh_court_location where loc_code in ('001', '002', '003');
DELETE FROM juror_mod.court_location where loc_code in ('001', '002', '003');

-- set app settings
INSERT INTO juror_mod.APP_SETTING (SETTING, VALUE) VALUES ('SLA_OVERDUE_DAYS', '10');
INSERT INTO juror_mod.APP_SETTING (SETTING, VALUE) VALUES ('URGENCY_DAYS', '10');


-- expenses things
DELETE FROM juror_mod.app_setting where setting = 'PAYMENT_AUTH_CODE';
DELETE FROM juror_mod.payment_data;

INSERT INTO juror_mod.app_setting (setting, value) VALUES ('PAYMENT_AUTH_CODE', 'testValue');

INSERT INTO juror_mod.rev_info (revision_number, revision_timestamp, changed_by) VALUES(0, 1713272656536, NULL);

INSERT INTO juror_mod.court_location_audit (revision, rev_type, loc_code, public_transport_soft_limit, taxi_soft_limit)
VALUES (0, 0, '415', 10.00000, 10.00000);

INSERT INTO juror_mod.expense_rates
(id, rate_per_mile_car_0_passengers, rate_per_mile_car_1_passengers, rate_per_mile_car_2_or_more_passengers, rate_per_mile_motorcycle_0_passengers, rate_per_mile_motorcycle_1_or_more_passengers, rate_per_mile_bike, limit_financial_loss_half_day, limit_financial_loss_full_day, limit_financial_loss_half_day_long_trial, limit_financial_loss_full_day_long_trial, rate_subsistence_standard, rate_subsistence_long_day)
VALUES(2, 0.31400, 0.35600, 0.39800, 0.31400, 0.32400, 0.09600, 40.00000, 64.95000, 64.95000, 129.91000, 5.71000, 12.17000);
INSERT INTO juror_mod.expense_rates
(id, rate_per_mile_car_0_passengers, rate_per_mile_car_1_passengers, rate_per_mile_car_2_or_more_passengers, rate_per_mile_motorcycle_0_passengers, rate_per_mile_motorcycle_1_or_more_passengers, rate_per_mile_bike, limit_financial_loss_half_day, limit_financial_loss_full_day, limit_financial_loss_half_day_long_trial, limit_financial_loss_full_day_long_trial, rate_subsistence_standard, rate_subsistence_long_day)
VALUES(3, 0.31400, 0.35600, 0.39800, 0.31400, 0.32400, 0.09600, 40.00000, 64.95000, 64.95000, 129.91000, 5.71000, 12.17000);
INSERT INTO juror_mod.expense_rates
(id, rate_per_mile_car_0_passengers, rate_per_mile_car_1_passengers, rate_per_mile_car_2_or_more_passengers, rate_per_mile_motorcycle_0_passengers, rate_per_mile_motorcycle_1_or_more_passengers, rate_per_mile_bike, limit_financial_loss_half_day, limit_financial_loss_full_day, limit_financial_loss_half_day_long_trial, limit_financial_loss_full_day_long_trial, rate_subsistence_standard, rate_subsistence_long_day)
VALUES(4, 0.31400, 0.35600, 0.39800, 0.31400, 0.32400, 0.09600, 32.47000, 40.00000, 64.95000, 129.91000, 5.71000, 12.17000);
INSERT INTO juror_mod.expense_rates
(id, rate_per_mile_car_0_passengers, rate_per_mile_car_1_passengers, rate_per_mile_car_2_or_more_passengers, rate_per_mile_motorcycle_0_passengers, rate_per_mile_motorcycle_1_or_more_passengers, rate_per_mile_bike, limit_financial_loss_half_day, limit_financial_loss_full_day, limit_financial_loss_half_day_long_trial, limit_financial_loss_full_day_long_trial, rate_subsistence_standard, rate_subsistence_long_day)
VALUES(5, 0.31500, 0.35600, 0.39800, 0.31400, 0.32400, 0.09600, 32.47000, 64.95000, 64.95000, 129.91000, 5.71000, 12.17000);
INSERT INTO juror_mod.expense_rates
(id, rate_per_mile_car_0_passengers, rate_per_mile_car_1_passengers, rate_per_mile_car_2_or_more_passengers, rate_per_mile_motorcycle_0_passengers, rate_per_mile_motorcycle_1_or_more_passengers, rate_per_mile_bike, limit_financial_loss_half_day, limit_financial_loss_full_day, limit_financial_loss_half_day_long_trial, limit_financial_loss_full_day_long_trial, rate_subsistence_standard, rate_subsistence_long_day)
VALUES(6, 0.31400, 0.35600, 0.39800, 0.31400, 0.32400, 0.09600, 40.00000, 64.95000, 64.95000, 129.91000, 5.71000, 12.17000);
INSERT INTO juror_mod.expense_rates
(id, rate_per_mile_car_0_passengers, rate_per_mile_car_1_passengers, rate_per_mile_car_2_or_more_passengers, rate_per_mile_motorcycle_0_passengers, rate_per_mile_motorcycle_1_or_more_passengers, rate_per_mile_bike, limit_financial_loss_half_day, limit_financial_loss_full_day, limit_financial_loss_half_day_long_trial, limit_financial_loss_full_day_long_trial, rate_subsistence_standard, rate_subsistence_long_day)
VALUES(7, 0.31400, 0.35600, 0.39800, 0.31400, 0.32400, 0.09600, 20.00000, 64.95000, 64.95000, 129.91000, 5.71000, 12.17000);
INSERT INTO juror_mod.expense_rates
(id, rate_per_mile_car_0_passengers, rate_per_mile_car_1_passengers, rate_per_mile_car_2_or_more_passengers, rate_per_mile_motorcycle_0_passengers, rate_per_mile_motorcycle_1_or_more_passengers, rate_per_mile_bike, limit_financial_loss_half_day, limit_financial_loss_full_day, limit_financial_loss_half_day_long_trial, limit_financial_loss_full_day_long_trial, rate_subsistence_standard, rate_subsistence_long_day)
VALUES(8, 0.31400, 0.35600, 0.39800, 0.31400, 0.32400, 0.09600, 50.00000, 64.95000, 64.95000, 129.91000, 5.71000, 12.17000);
INSERT INTO juror_mod.expense_rates
(id, rate_per_mile_car_0_passengers, rate_per_mile_car_1_passengers, rate_per_mile_car_2_or_more_passengers, rate_per_mile_motorcycle_0_passengers, rate_per_mile_motorcycle_1_or_more_passengers, rate_per_mile_bike, limit_financial_loss_half_day, limit_financial_loss_full_day, limit_financial_loss_half_day_long_trial, limit_financial_loss_full_day_long_trial, rate_subsistence_standard, rate_subsistence_long_day)
VALUES(9, 0.31400, 0.35600, 0.39800, 0.31400, 0.32400, 0.09600, 50.00000, 64.95000, 64.95000, 129.91000, 5.71000, 12.17000);
INSERT INTO juror_mod.expense_rates
(id, rate_per_mile_car_0_passengers, rate_per_mile_car_1_passengers, rate_per_mile_car_2_or_more_passengers, rate_per_mile_motorcycle_0_passengers, rate_per_mile_motorcycle_1_or_more_passengers, rate_per_mile_bike, limit_financial_loss_half_day, limit_financial_loss_full_day, limit_financial_loss_half_day_long_trial, limit_financial_loss_full_day_long_trial, rate_subsistence_standard, rate_subsistence_long_day)
VALUES(10, 0.31400, 0.35600, 0.39800, 0.31400, 0.32400, 0.09600, 70.00000, 64.95000, 64.95000, 129.91000, 5.71000, 12.17000);
INSERT INTO juror_mod.expense_rates
(id, rate_per_mile_car_0_passengers, rate_per_mile_car_1_passengers, rate_per_mile_car_2_or_more_passengers, rate_per_mile_motorcycle_0_passengers, rate_per_mile_motorcycle_1_or_more_passengers, rate_per_mile_bike, limit_financial_loss_half_day, limit_financial_loss_full_day, limit_financial_loss_half_day_long_trial, limit_financial_loss_full_day_long_trial, rate_subsistence_standard, rate_subsistence_long_day)
VALUES(11, 0.31400, 0.35600, 0.39800, 0.31400, 0.32400, 0.09600, 90.00000, 64.95000, 64.95000, 129.91000, 5.71000, 12.17000);
INSERT INTO juror_mod.expense_rates
(id, rate_per_mile_car_0_passengers, rate_per_mile_car_1_passengers, rate_per_mile_car_2_or_more_passengers, rate_per_mile_motorcycle_0_passengers, rate_per_mile_motorcycle_1_or_more_passengers, rate_per_mile_bike, limit_financial_loss_half_day, limit_financial_loss_full_day, limit_financial_loss_half_day_long_trial, limit_financial_loss_full_day_long_trial, rate_subsistence_standard, rate_subsistence_long_day)
VALUES(12, 0.31400, 0.35600, 0.39800, 0.31400, 0.32400, 0.09600, 37.45000, 64.95000, 64.95000, 129.91000, 5.71000, 12.17000);
INSERT INTO juror_mod.expense_rates
(id, rate_per_mile_car_0_passengers, rate_per_mile_car_1_passengers, rate_per_mile_car_2_or_more_passengers, rate_per_mile_motorcycle_0_passengers, rate_per_mile_motorcycle_1_or_more_passengers, rate_per_mile_bike, limit_financial_loss_half_day, limit_financial_loss_full_day, limit_financial_loss_half_day_long_trial, limit_financial_loss_full_day_long_trial, rate_subsistence_standard, rate_subsistence_long_day)
VALUES(13, 0.31400, 0.35600, 0.39800, 0.31400, 0.32400, 0.09600, 38.45000, 65.95000, 64.95000, 129.91000, 5.71000, 12.17000);
INSERT INTO juror_mod.expense_rates
(id, rate_per_mile_car_0_passengers, rate_per_mile_car_1_passengers, rate_per_mile_car_2_or_more_passengers, rate_per_mile_motorcycle_0_passengers, rate_per_mile_motorcycle_1_or_more_passengers, rate_per_mile_bike, limit_financial_loss_half_day, limit_financial_loss_full_day, limit_financial_loss_half_day_long_trial, limit_financial_loss_full_day_long_trial, rate_subsistence_standard, rate_subsistence_long_day)
VALUES(14, 0.31400, 0.35600, 0.39800, 0.31400, 0.32400, 0.09600, 37.45000, 64.95000, 64.95000, 129.91000, 5.71000, 12.17000);
INSERT INTO juror_mod.expense_rates
(id, rate_per_mile_car_0_passengers, rate_per_mile_car_1_passengers, rate_per_mile_car_2_or_more_passengers, rate_per_mile_motorcycle_0_passengers, rate_per_mile_motorcycle_1_or_more_passengers, rate_per_mile_bike, limit_financial_loss_half_day, limit_financial_loss_full_day, limit_financial_loss_half_day_long_trial, limit_financial_loss_full_day_long_trial, rate_subsistence_standard, rate_subsistence_long_day)
VALUES(15, 0.31400, 0.35600, 0.39800, 0.31400, 0.32400, 0.09600, 50.00000, 64.95000, 64.95000, 129.91000, 5.71000, 12.17000);
INSERT INTO juror_mod.expense_rates
(id, rate_per_mile_car_0_passengers, rate_per_mile_car_1_passengers, rate_per_mile_car_2_or_more_passengers, rate_per_mile_motorcycle_0_passengers, rate_per_mile_motorcycle_1_or_more_passengers, rate_per_mile_bike, limit_financial_loss_half_day, limit_financial_loss_full_day, limit_financial_loss_half_day_long_trial, limit_financial_loss_full_day_long_trial, rate_subsistence_standard, rate_subsistence_long_day)
VALUES(16, 0.31400, 0.35600, 0.39800, 0.31400, 0.32400, 0.09600, 37.45000, 64.95000, 64.95000, 129.91000, 5.71000, 12.17000);
