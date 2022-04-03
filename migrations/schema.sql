create table barcode_mask
(
	code varchar(255) not null
		constraint barcode_mask_pkey
			primary key,
	external_id varchar(255),
	name_ru varchar(255),
	name_kz varchar(255),
	name_en varchar(255)
);


create table app_user
(
	id bigint not null
		constraint app_user_pkey
			primary key,
	iin varchar(255),
	first_name varchar(255),
	middle_name varchar(255),
	last_name varchar(255),
	phone varchar(255)
		constraint phone_unique_constraint
			unique,
	username varchar(255)
		constraint username_unique_constraint
			unique,
	email varchar(255)
		constraint email_unique_constraint
			unique,
	password varchar(255),
	is_active boolean,
	active_directory_link varchar(255),
	reg_datetime timestamp,
	otdel_id bigint
		constraint fk_app_user__otdel_id
			references otdel,
	organization_id bigint
		constraint fk_app_user__organization_id
			references organization
);


create table app_role
(
	id bigint not null
		constraint app_role_pkey
			primary key,
	code varchar(255) not null,
	name_ru varchar(255),
	name_kz varchar(255),
	name_en varchar(255)
);


create table category
(
	id bigint not null
		constraint category_pkey
			primary key,
	code varchar(255),
	external_id varchar(255),
	name_ru varchar(255)
		constraint ux_category_name_ru
			unique,
	name_kz varchar(255),
	name_en varchar(255)
);


create table certificate_type
(
	id bigint not null
		constraint certificate_type_pkey
			primary key,
	external_id varchar(255),
	name_ru varchar(255),
	name_kz varchar(255),
	name_en varchar(255)
);


create table child_weight
(
	id bigint not null
		constraint child_weight_pkey
			primary key,
	external_id varchar(255),
	name_ru varchar(255),
	name_kz varchar(255),
	name_en varchar(255)
);


create table color
(
	id bigint not null
		constraint color_pkey
			primary key,
	external_id varchar(255),
	name_ru varchar(255),
	name_kz varchar(255),
	name_en varchar(255)
);


create table commercial_conditions
(
	id bigint not null
		constraint commercial_conditions_pkey
			primary key,
	fix_bonus numeric(21,2),
	fix_premium numeric(21,2),
	logist_cost numeric(21,2),
	logist_bonus numeric(21,2),
	payment_due_date varchar(255),
	account_work varchar(255),
	inkoterms varchar(255),
	minimal_order_lot varchar(255),
	status varchar(255),
	payment_type_id bigint
		constraint fk_commercial_conditions__payment_type_id
			references payment_type,
	organization_id bigint
		constraint fk_commercial_conditions__organization_id
			references organization
);


create table commercial_proposal
(
	id bigint not null
		constraint commercial_proposal_pkey
			primary key,
	process_instance varchar(255),
	process_date timestamp,
	status varchar(255),
	proposal_conditions_id bigint
		constraint ux_commercial_proposal__proposal_conditions_id
			unique
		constraint fk_commercial_proposal__proposal_conditions_id
			references proposal_conditions,
	organization_id bigint
		constraint fk_commercial_proposal__organization_id
			references organization,
	otdel_id bigint
		constraint fk_commercial_proposal__otdel_id
			references otdel,
	delivery_conditions_id bigint
		constraint fk_commercial_proposal__delivery_conditions_id
			references delivery_conditions,
	enter_product_request_id bigint
		constraint fk_commercial_proposal__enter_product_request_id
			references enter_product_request,
	currency varchar(10)
);


create table commercial_proposal_status
(
	code varchar(255) not null
		constraint commercial_proposal_status_pkey
			primary key,
	name_ru varchar(255),
	name_kz varchar(255),
	name_en varchar(255)
);


create table commercial_proposal_status_h
(
	id bigint not null
		constraint commercial_proposal_status_h_pkey
			primary key,
	status_date timestamp,
	status varchar(255),
	comment_text text,
	user_id bigint
		constraint fk_commercial_proposal_status_h__user_id
			references app_user,
	commercial_proposal_id bigint
		constraint fk_commercial_proposal_status_h__commercial_proposal_id
			references commercial_proposal,
	is_internal boolean default false
);


create table country
(
	id bigint not null
		constraint country_pkey
			primary key,
	external_id varchar(255),
	name_ru varchar(255),
	name_kz varchar(255),
	name_en varchar(255)
);


create table currency
(
	code varchar(255) not null
		constraint currency_pkey
			primary key,
	name_ru varchar(255),
	name_kz varchar(255),
	name_en varchar(255)
);

create table delivery_conditions
(
	id bigint not null
		constraint delivery_conditions_pkey
			primary key,
	external_id varchar(255),
	name_ru varchar(255),
	name_kz varchar(255),
	name_en varchar(255)
);

create table enter_product_request
(
	id bigint not null
		constraint enter_product_request_pkey
			primary key,
	start_date timestamp,
	end_date timestamp,
	status varchar(255),
	otdel_id bigint
		constraint fk_enter_product_request__otdel_id
			references otdel,
	organization_id bigint
		constraint fk_enter_product_request__organization_id
			references organization
);

create table enter_product_request_notif
(
	id bigint not null
		constraint enter_product_request_notif_pkey
			primary key,
	create_date timestamp,
	text varchar(255),
	enter_product_request_id bigint
		constraint fk_enter_product_request_notif__enter_product_request_id
			references enter_product_request
);

create table enter_product_request_status
(
	code varchar(255) not null
		constraint enter_product_request_status_pkey
			primary key,
	name_ru varchar(255),
	name_kz varchar(255),
	name_en varchar(255)
);


create table enter_product_request_status_h
(
	id bigint not null
		constraint enter_product_request_status_h_pkey
			primary key,
	status varchar(255),
	status_date_time timestamp,
	comment_text varchar(255),
	user_id bigint
		constraint fk_enter_product_request_status_h__user_id
			references app_user,
	enter_product_request_id bigint
		constraint fk_enter_product_request_status_h__enter_product_request_id
			references enter_product_request,
	is_internal boolean default false
);


create table fat_content
(
	id bigint not null
		constraint fat_content_pkey
			primary key,
	external_id varchar(255),
	name_ru varchar(255),
	name_kz varchar(255),
	name_en varchar(255)
);


create table field_group
(
	id bigint not null
		constraint field_group_pkey
			primary key,
	external_id varchar(255),
	name_ru varchar(255),
	name_kz varchar(255),
	name_en varchar(255)
);


create table foliage
(
	id bigint not null
		constraint foliage_pkey
			primary key,
	external_id varchar(255),
	name_ru varchar(255),
	name_kz varchar(255),
	name_en varchar(255)
);


create table form_of_law
(
	id bigint not null
		constraint form_of_law_pkey
			primary key,
	code varchar(255),
	is_resident boolean,
	name_ru varchar(255),
	name_kz varchar(255),
	name_en varchar(255)
);


create table jhi_group
(
	id bigint not null
		constraint jhi_group_pkey
			primary key,
	code varchar(255),
	external_id varchar(255),
	name_ru varchar(255),
	name_kz varchar(255),
	name_en varchar(255),
	otdel_id bigint
);


create table kbe
(
	id bigint not null
		constraint kbe_pkey
			primary key,
	code varchar(255),
	name_ru varchar(255),
	name_kz varchar(255),
	name_en varchar(255)
);


create table load
(
	id bigint not null
		constraint load_pkey
			primary key,
	external_id varchar(255),
	name_ru varchar(255),
	name_kz varchar(255),
	name_en varchar(255)
);


create table logisticts_options
(
	id bigint not null
		constraint logisticts_options_pkey
			primary key,
	external_id varchar(255),
	name_ru varchar(255),
	name_kz varchar(255),
	name_en varchar(255)
);


create table measurement
(
	id bigint not null
		constraint measurement_pkey
			primary key,
	external_id varchar(255),
	name_ru varchar(255),
	name_kz varchar(255),
	name_en varchar(255)
);


create table organization
(
	id bigint not null
		constraint organization_pkey
			primary key,
	bin varchar(255)
		constraint bin_unique_constraint
			unique,
	name varchar(255),
	is_resident boolean,
	corp_phone varchar(255),
	contact_fio varchar(255),
	website varchar(255),
	gln varchar(255)
		constraint gln_unique_constraint
			unique,
	is_stm boolean,
	is_nds boolean,
	iban varchar(255)
		constraint iban_unique_constraint
			unique,
	bic varchar(255),
	kbe_code varchar(255),
	address varchar(255),
	form_of_law_id bigint
		constraint fk_organization__form_of_law_id
			references form_of_law,
	kbe_id bigint
		constraint fk_organization__kbe_id
			references kbe
);


create table department
(
	id bigint not null
		constraint otdel_pkey
			primary key,
	code varchar(255),
	external_id varchar(255),
	name_ru varchar(255),
	name_kz varchar(255),
	name_en varchar(255),
	level varchar(255)
);

create table package_type
(
	id bigint not null
		constraint package_type_pkey
			primary key,
	external_id varchar(255),
	short_name_ru varchar(255),
	short_name_kz varchar(255),
	short_name_en varchar(255),
	name_ru varchar(255),
	name_kz varchar(255),
	name_en varchar(255)
);

alter table package_type owner to suppliers_portal;

create table packaging
(
	id bigint not null
		constraint packaging_pkey
			primary key,
	external_id varchar(255),
	name_ru varchar(255),
	name_kz varchar(255),
	name_en varchar(255)
);

alter table packaging owner to suppliers_portal;

create table payment_type
(
	id bigint not null
		constraint payment_type_pkey
			primary key,
	external_id varchar(255),
	name_ru varchar(255),
	name_kz varchar(255),
	name_en varchar(255)
);

alter table payment_type owner to suppliers_portal;

create table photo_type
(
	id bigint not null
		constraint photo_type_pkey
			primary key,
	name_ru varchar(255),
	name_kz varchar(255),
	name_en varchar(255)
);

alter table photo_type owner to suppliers_portal;


create table product
(
	id bigint not null
		constraint product_pkey
			primary key,
	gtin_code varchar(255),
	name_ru varchar(255),
	name_kz varchar(255),
	name_en varchar(255),
	stage varchar(255),
	item_name varchar(255),
	sku_features_for_item_name varchar(255),
	item_price_tag_name varchar(255),
	supplier varchar(255),
	trademark varchar(255),
	brand varchar(255),
	manufacturer varchar(255),
	composition_kazakh_lang text,
	composition_russian_lang text,
	assign_barcode boolean,
	vendor_code varchar(255),
	material varchar(255),
	sale_period_days varchar(255),
	expiration_date_after_opening_days varchar(255),
	density varchar(255),
	watt varchar(255),
	ampere varchar(255),
	plinth varchar(255),
	model varchar(255),
	completeness varchar(255),
	halal_sign boolean,
	lactose_free_sign boolean,
	gluten_free_sign boolean,
	sugar_free_sign boolean,
	kosher_sign boolean,
	organic_sign boolean,
	vegan_sign boolean,
	diabetic_sign boolean,
	natural_loss_rate varchar(255),
	proteins varchar(255),
	fats varchar(255),
	carbohydrates varchar(255),
	energy_value_kcal varchar(255),
	touch_sensitivity_sign boolean,
	ethylene_sensitivity boolean,
	alcohol_strength varchar(255),
	tare_weight varchar(255),
	length_deep varchar(255),
	width varchar(255),
	height varchar(255),
	net_weight varchar(255),
	gross_weight varchar(255),
	dry_weight varchar(255),
	minimum_storage_temperature varchar(255),
	maximum_storage_temperature varchar(255),
	block_attribute boolean,
	wblock_code varchar(255),
	number_of_pieceskg_in_block varchar(255),
	length_depth_cm varchar(255),
	width_cm varchar(255),
	height_cm varchar(255),
	weight_cm varchar(255),
	unit_gross_weight_kg varchar(255),
	conformity_certificate_end_date timestamp,
	position_name_in_state_language varchar(255),
	price_segment varchar(255),
	sale_price_with_va_ttg_cc varchar(255),
	sale_price_with_va_ttg_atak varchar(255),
	pin_code varchar(255),
	packaging numeric(21,2),
	load numeric(21,2),
	foliage varchar(255),
	fat_content numeric(21,2),
	logisticts_options varchar(255),
	sizeval varchar(255),
	additional_options varchar(255),
	age_category integer,
	barcode_mask varchar(255),
	origin_country_id bigint
		constraint fk_product__origin_country_id
			references country,
	tnvd_id bigint
		constraint fk_product__tnvd_id
			references tnvd,
	color_id bigint
		constraint fk_product__color_id
			references color,
	taste_id bigint
		constraint fk_product__taste_id
			references taste,
	package_type_id bigint
		constraint fk_product__package_type_id
			references package_type,
	child_weight_id bigint
		constraint fk_product__child_weight_id
			references child_weight,
	storage_and_transportation_conditions_id bigint
		constraint fk_product__storage_and_transportation_conditions_id
			references storage_and_transportation_conditions,
	delivery_conditions_id bigint
		constraint fk_product__delivery_conditions_id
			references delivery_conditions,
	measurement_id bigint
		constraint fk_product__measurement_id
			references measurement,
	organization_id bigint
		constraint fk_product__organization_id
			references organization,
	sub_category_id bigint
		constraint fk_product__sub_category_id
			references sub_category,
	subgroup_id bigint
		constraint fk_product__subgroup_id
			references subgroup,
	item_photo_with_line bigint,
	item_photo_with_front_side bigint,
	composition_photo bigint,
	item_photo_for_scales bigint,
	certificate_photo_of_item_conformity bigint,
	constraint product_gtin_code_organization_id_key
		unique (gtin_code, organization_id),
	constraint product_vendor_code_organization_id_key
		unique (vendor_code, organization_id)
);


create table product_certificate
(
	id bigint not null
		constraint product_certificate_pkey
			primary key,
	name varchar(255),
	active boolean,
	filename varchar(255),
	filepath varchar(255),
	product_id bigint
		constraint fk_product_certificate__product_id
			references product,
	content_type varchar(255) not null,
	original_filename varchar(500) not null,
	expire_date date,
	certificate_type_id bigint
		constraint fk_certificate
			references certificate_type,
	url varchar(500)
);


create table product_field
(
	id bigint not null
		constraint product_field_pkey
			primary key,
	code varchar(255),
	sprut_code varchar(255),
	mandatory boolean,
	field_name varchar(255),
	name_ru varchar(255),
	name_kz varchar(255),
	name_en varchar(255),
	field_type varchar(255),
	field_group_id bigint
		constraint fk_product_field__field_group_id
			references field_group
);


create table product_field_in_category
(
	id bigint not null
		constraint product_field_in_category_pkey
			primary key,
	active boolean,
	sub_category_id bigint
		constraint fk_product_field_in_category__sub_category_id
			references sub_category,
	product_field_id bigint
		constraint fk_product_field_in_category__product_field_id
			references product_field,
	constraint product_field_category_unique_constraint
		unique (sub_category_id, product_field_id),
	constraint product_field_in_category_sub_category_id_product_field_id_key
		unique (sub_category_id, product_field_id)
);


create table product_field_stage
(
	id bigint not null
		constraint product_field_stage_pkey
			primary key,
	stage varchar(255),
	privilege_supplier varchar(255),
	privilege_km varchar(255),
	privilege_rn varchar(255),
	privilege_oukd varchar(255),
	privilege_ducp varchar(255),
	product_field_id bigint
		constraint fk_product_field_stage__product_field_id
			references product_field
);


create table product_instance
(
	id bigint not null
		constraint product_instance_pkey
			primary key,
	status_change_date timestamp,
	price_excluding_vat numeric(21,2),
	vat_rate numeric(21,2),
	price_with_vat numeric(21,2),
	discount_from_basic_price numeric(21,2),
	marginal_trade_markup numeric(21,2),
	status varchar(255),
	commercial_proposal_id bigint
		constraint fk_product_instance__commercial_proposal_id
			references commercial_proposal,
	product_id bigint
		constraint fk_product_instance__product_id
			references product,
	region_id bigint
		constraint fk_product_instance__region_id
			references region,
	enter_product_request_id bigint
		constraint fk_product_instance__enter_product_request_id
			references enter_product_request
);


create table product_photo
(
	id bigint not null
		constraint product_photo_pkey
			primary key,
	input_dir varchar(255),
	filename varchar(255),
	original_filename varchar(255),
	url varchar(255),
	photo_type_id bigint,
	product_id bigint,
	content_type varchar(255) not null,
	filepath varchar(500) not null
);


create table proposal_conditions
(
	id bigint not null
		constraint proposal_conditions_pkey
			primary key,
	fix_bonus numeric(21,2),
	fix_premium numeric(21,2),
	logist_bonus numeric(21,2),
	account_work varchar(255),
	inkoterms varchar(255),
	minimal_order_lot varchar(255),
	payment_due_date varchar(50),
	payment_type_id bigint
		constraint fk_proposal_conditions_payment_type
			references payment_type
);


create table region
(
	id bigint not null
		constraint region_pkey
			primary key,
	code varchar(255),
	external_id varchar(255),
	name_ru varchar(255),
	name_kz varchar(255),
	name_en varchar(255)
);


create table rel_app_user__role
(
	role_id bigint not null
		constraint fk_rel_app_user__role__role_id
			references app_role,
	app_user_id bigint not null
		constraint fk_rel_app_user__role__app_user_id
			references app_user,
	constraint rel_app_user__role_pkey
		primary key (app_user_id, role_id)
);


create table rel_category__role
(
	role_id bigint not null
		constraint fk_rel_category__role__role_id
			references app_role,
	category_id bigint not null
		constraint fk_rel_category__role__category_id
			references category,
	constraint rel_category__role_pkey
		primary key (category_id, role_id)
);


create table rel_commercial_conditions__region
(
	region_id bigint not null
		constraint fk_rel_commercial_conditions__region__region_id
			references region,
	commercial_conditions_id bigint not null
		constraint fk_rel_commercial_conditions__region__commercial_conditions_id
			references commercial_conditions,
	constraint rel_commercial_conditions__region_pkey
		primary key (commercial_conditions_id, region_id)
);

create table rel_organization__region
(
	region_id bigint not null
		constraint fk_rel_organization__region__region_id
			references region,
	organization_id bigint not null
		constraint fk_rel_organization__region__organization_id
			references organization,
	constraint rel_organization__region_pkey
		primary key (organization_id, region_id)
);

alter table rel_organization__region owner to suppliers_portal;

create table rel_proposal_conditions__region
(
	region_id bigint not null
		constraint fk_rel_proposal_conditions__region__region_id
			references region,
	proposal_conditions_id bigint not null
		constraint fk_rel_proposal_conditions__region__proposal_conditions_id
			references proposal_conditions,
	constraint rel_proposal_conditions__region_pkey
		primary key (proposal_conditions_id, region_id)
);

alter table rel_proposal_conditions__region owner to suppliers_portal;

create table request_step_notification
(
	id bigint not null
		constraint request_step_notification_pkey
			primary key,
	create_date timestamp,
	bprocess varchar(255),
	jhi_to varchar(255),
	mail_text varchar(255),
	subject varchar(255),
	json varchar(255),
	user_id bigint
		constraint fk_request_step_notification__user_id
			references app_user
);

create table reset_password_application
(
	id bigint not null
		constraint reset_password_application_pkey
			primary key,
	email varchar(100),
	recovery_code varchar(8),
	created_time timestamp,
	token varchar(255),
	confirmation_status boolean
);

alter table reset_password_application owner to suppliers_portal;

create table storage_and_transportation_conditions
(
	id bigint not null
		constraint storage_and_transportation_conditions_pkey
			primary key,
	external_id varchar(255),
	name_ru varchar(255),
	name_kz varchar(255),
	name_en varchar(255)
);

create table sub_category
(
	id bigint not null
		constraint sub_category_pkey
			primary key,
	code varchar(255),
	external_id varchar(255),
	name_ru varchar(255),
	name_kz varchar(255),
	name_en varchar(255),
	otdel_id bigint
		constraint fk_sub_category__otdel_id
			references otdel,
	category_id bigint
		constraint fk_sub_category__category_id
			references category,
	constraint ux_subcategory_category_name_ru
		unique (category_id, name_ru)
);

create table subgroup
(
	id bigint not null
		constraint subgroup_pkey
			primary key,
	code varchar(255),
	external_id varchar(255),
	name_ru varchar(255),
	name_kz varchar(255),
	name_en varchar(255),
	group_id bigint
);

create table tnvd
(
	id bigint not null
		constraint tnvd_pkey
			primary key,
	external_id varchar(255),
	code varchar(255),
	name_ru varchar(1500),
	name_kz varchar(1500),
	name_en varchar(1500)
);


