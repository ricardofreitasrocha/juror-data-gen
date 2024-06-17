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

INSERT INTO juror_mod.expense_rates(id, rate_per_mile_car_0_passengers, rate_per_mile_car_1_passengers,
                                    rate_per_mile_car_2_or_more_passengers, rate_per_mile_motorcycle_0_passengers,
                                    rate_per_mile_motorcycle_1_or_more_passengers, rate_per_mile_bike,
                                    limit_financial_loss_half_day, limit_financial_loss_full_day,
                                    limit_financial_loss_half_day_long_trial, limit_financial_loss_full_day_long_trial,
                                    rate_subsistence_standard, rate_subsistence_long_day)
values (1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1),
       (2, 0.314, 0.356, 0.398, 0.314, 0.324, 0.096, 32.47, 64.95, 64.95, 129.91, 5.71, 12.17);
