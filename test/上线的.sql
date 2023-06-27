
cfg: "root:?cr12alD0KLT@tcp(147.139.185.188:3306)/rupiah_loan?timeout=3s&parseTime=true&loc=Local&charset=utf8" 		#api 库

DROP TABLE IF EXISTS `google`;
CREATE TABLE `google` (
	`id` bigint(20) NOT NULL AUTO_INCREMENT,
	`uid` bigint(20) DEFAULT NULL COMMENT '用户ID',
	`app_instance_id` char(80) DEFAULT NULL COMMENT '谷歌ID',
	`created` bigint(20) CURRENT_TIMESTAMP DEFAULT NULL,
	`updated` bigint(20) CURRENT_TIMESTAMP DEFAULT NULL,
	PRIMARY KEY (`id`),
	KEY `idx_uid` (`uid`),
	KEY `idx_app_instance_id` (`app_instance_id`)
	) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='Google 事件app上传app_instance_id记录表'

ALTER TABLE `rupiah_loan`.`google` 
MODIFY COLUMN `app_instance_id` char(80) DEFAULT NULL COMMENT '谷歌ID' AFTER `uid`



CREATE TABLE `google_event` (
	`id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id',
	`uid` bigint(20) DEFAULT NULL COMMENT '用户ID',
	`app_instance_id` char(80) NOT NULL DEFAULT '' COMMENT '谷歌ID',
	`event_name` varchar(50) NOT NULL DEFAULT '' COMMENT '事件名称',
	`event_value` varchar(50) NOT NULL DEFAULT '' COMMENT '事件值',
	`event_time` varchar(20) NOT NULL DEFAULT '' COMMENT '事件时间',
	`response` varchar(255) NOT NULL DEFAULT '0' COMMENT '响应结果',
	`created` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建时间',
	`updated` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建时间',
	PRIMARY KEY (`id`),
	KEY `idx_uid` (`uid`) USING BTREE,
	KEY `idx_created` (`created`) USING BTREE,
	KEY `idx_app_instance_id` (`app_instance_id`) USING BTREE
  ) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='Google 事件打点记录表';

  ALTER TABLE `rupiah_loan`.`google_event` 
MODIFY COLUMN `app_instance_id` char(80) NOT NULL DEFAULT 1 COMMENT '谷歌ID' AFTER `uid`

==========================================================================================================================================================================================
上线的sql

CREATE TABLE `google` (
  `id` bigint(20)  UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `uid` bigint(20) DEFAULT NULL COMMENT '用户ID',
  `app_instance_id` varchar(80) DEFAULT NULL COMMENT '谷歌ID',
  `created` bigint(20) DEFAULT NULL COMMENT '创建时间',
  `updated` bigint(20) DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `idx_uid` (`uid`),
  KEY `idx_app_instance_id` (`app_instance_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='Google 事件app上传app_instance_id记录表';

CREATE TABLE `google_event` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `uid` bigint(20) DEFAULT NULL COMMENT '用户ID',
  `app_instance_id` varchar(80) NOT NULL DEFAULT '1' COMMENT '谷歌ID',
  `event_name` varchar(50) NOT NULL DEFAULT '' COMMENT '事件名称',
  `event_value` varchar(50) NOT NULL DEFAULT '' COMMENT '事件值',
  `event_time` varchar(20) NOT NULL DEFAULT '' COMMENT '事件时间',
  `response` varchar(255) NOT NULL DEFAULT '0' COMMENT '响应结果',
  `created` bigint(20) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated` bigint(20) NOT NULL DEFAULT '0' COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `idx_uid` (`uid`) USING BTREE,
  KEY `idx_created` (`created`) USING BTREE,
  KEY `idx_app_instance_id` (`app_instance_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='Google 事件打点记录表';

==========================================================================================================================================================================================
// IziOCR需求

// 新增mysql连接
-- 第一版

CREATE TABLE `third_ocr` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `img_frist` varchar(120) NOT NULL DEFAULT '' COMMENT '第一张图片前100位',
  `img_second` varchar(120) NOT NULL DEFAULT '' COMMENT '第二张图片前100位',
  `type` int(10) NOT NULL DEFAULT 0 COMMENT '类型:1 advcance 2 izi',
  `created` bigint(20) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated` bigint(20) NOT NULL DEFAULT '0' COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `idx_img_frist` (`img_frist`) USING BTREE,
  KEY `idx_img_second` (`img_second`) USING BTREE,
  KEY `idx_created` (`created`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='izi_ocr 身份证认证记录表';

CREATE TABLE `third_res` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `req` varchar(120) DEFAULT NULL COMMENT '请求传参',
  `code` varchar(20) NOT NULL DEFAULT '1' COMMENT '返回code',
  `msg` varchar(255) NOT NULL DEFAULT '' COMMENT '返回msg',
  `req_url` varchar(255) NOT NULL DEFAULT '' COMMENT '请求url',
  `extra` varchar(255) NOT NULL DEFAULT '' COMMENT '其他',
  `response` text NOT NULL  COMMENT '返回结果',
  `created` bigint(20) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated` bigint(20) NOT NULL DEFAULT '0' COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `idx_code` (`code`) USING BTREE,
  KEY `idx_created` (`created`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='third_res 第三方结果返回';

-- 第二版

CREATE TABLE `third_req_res_record` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `data_type` int(10) NOT NULL DEFAULT 0 COMMENT '数据类型:1 req 2 res',
  `img_frist` varchar(120) NOT NULL DEFAULT '' COMMENT '第一张图片后30位',
  `img_second` varchar(120) NOT NULL DEFAULT '' COMMENT '第二张图片后30位',
  `type` int(10) NOT NULL DEFAULT 0 COMMENT '类型:1 advcance 2 izi',
  `type_tag` int(10) NOT NULL DEFAULT 0 COMMENT '类型标签:1 身份证OCR 2 人脸比对',
  `req` varchar(120) DEFAULT NULL COMMENT '请求传参',
  `code` varchar(20) NOT NULL DEFAULT '1' COMMENT '返回code',
  `msg` varchar(255) NOT NULL DEFAULT '' COMMENT '返回msg',
  `req_url` varchar(255) NOT NULL DEFAULT '' COMMENT '请求url',
  `extra` varchar(255) NOT NULL DEFAULT '' COMMENT '其他',
  `response` text  COMMENT '返回结果',
  `created` bigint(20) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated` bigint(20) NOT NULL DEFAULT '0' COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `idx_img_frist` (`img_frist`) USING BTREE,
  KEY `idx_img_second` (`img_second`) USING BTREE,
  KEY `idx_created` (`created`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='第三方传参/返回记录表';


==========================================================================================================================================================================================

老用户冲量需求上线sql

ALTER TABLE `user_coupon` ADD `is_rid` INT  NULL  DEFAULT '1'  COMMENT '用户优惠券是否弹过 1没有弹过 2弹过'  AFTER `desc`;
ALTER TABLE `coupon_info` ADD `desc` TEXT  CHARACTER SET utf8  COLLATE utf8_general_ci  NULL  COMMENT '优惠券规则'  AFTER `action`;

656---9008
INSERT INTO `coupon_info` ( `name`, `action`, `desc`, `start_time`, `end_time`, `type`, `status`, `is_path`, `is_show`, `num_day`, `rollover_amount`, `interest_amount`, `poundage_amount`, `manage_amount`, `overdue_amount`, `late_fee_amount`, `money`, `all_money`, `min_money`, `mix_money`, `min_overdue_day`, `mix_overdue_day`, `loan_product_id`, `order_num`, `ctime`)
VALUES
	('biaya layanan', '', '1. Kupon ini hanya berlaku untuk pemotongan 2% dari biaya layanan dan bukan untuk bunga pinjaman<br>\\n2. Hanya berlaku untuk produk 90 hari<br>\\n3. Hanya berlaku selama 90 hari dan kadaluarsa jika lebih dari 90 hari. Mohon untuk bayar tepat waktu<br>\\n4. Voucher ini hanya berlaku untuk pembayaran cicilan pertama<br>\\n', 1590940800, 1952037600, 4, 1, 1, 1, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, '108', 0, 1681784833);

672--9009
INSERT INTO `coupon_info` (`name`, `action`, `desc`, `start_time`, `end_time`, `type`, `status`, `is_path`, `is_show`, `num_day`, `rollover_amount`, `interest_amount`, `poundage_amount`, `manage_amount`, `overdue_amount`, `late_fee_amount`, `money`, `all_money`, `min_money`, `mix_money`, `min_overdue_day`, `mix_overdue_day`, `loan_product_id`, `order_num`, `ctime`)
VALUES
	('biaya layanan', '', '1. Kupon ini hanya berlaku untuk pemotongan 2% dari biaya layanan dan bukan untuk bunga pinjaman<br>\\n2. Hanya berlaku untuk produk 60, 90 dan 120 hari<br>\\n3. Hanya berlaku selama 90 hari dan kadaluarsa jika lebih dari 90 hari. Mohon untuk bayar tepat waktu<br>\\n4. Voucher ini hanya berlaku untuk pembayaran cicilan pertama<br>\\n', 1590940800, 1952037600, 4, 1, 1, 1, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, '', 0, 1681784833);



/* 17:25:05 pro-uangme_loan uangme_loan */ UPDATE `coupon_info` SET `desc` = '1. Voucher 100 ribu rupiah ini dapat digunakan untuk diskon biaya pelayanan dan dapat digunakan dengan meminjam pinjaman<br>2. Hanya berlaku untuk produk 60, 90 dan 120 hari<br>3. Hanya berlaku selama 90 hari dan kadaluarsa jika lebih dari 90 hari. Mohon untuk bayar tepat waktu<br>4. Voucher ini hanya berlaku untuk pembayaran cicilan pertama<br>' WHERE `id` = '151';
/* 17:25:05 pro-uangme_loan uangme_loan */ UPDATE `coupon_info` SET `desc` = '1. Kupon ini hanya berlaku untuk pemotongan 2% dari biaya layanan dan bukan untuk bunga pinjaman<br>2. Hanya berlaku untuk produk 90 hari<br>3. Hanya berlaku selama 90 hari dan kadaluarsa jika lebih dari 90 hari. Mohon untuk bayar tepat waktu<br>4. Voucher ini hanya berlaku untuk pembayaran cicilan pertama<br>' WHERE `id` = '9008';
/* 17:25:05 pro-uangme_loan uangme_loan */ UPDATE `coupon_info` SET `desc` = '1. Kupon ini hanya berlaku untuk pemotongan 2% dari biaya layanan dan bukan untuk bunga pinjaman<br>2. Hanya berlaku untuk produk 60, 90 dan 120 hari<br>3. Hanya berlaku selama 90 hari dan kadaluarsa jika lebih dari 90 hari. Mohon untuk bayar tepat waktu<br>4. Voucher ini hanya berlaku untuk pembayaran cicilan pertama<br>' WHERE `id` = '9009';










阿波罗配置

					desc = viper.GetString("new_useractivation_coupons")


new_useractivation_coupons_rule:"1. Voucher 100 ribu rupiah ini dapat digunakan untuk diskon biaya pelayanan dan dapat digunakan dengan meminjam pinjaman<br>2. Hanya berlaku untuk produk 60, 90 dan 120 hari<br>3. Hanya berlaku selama 90 hari dan kadaluarsa jika lebih dari 90 hari. Mohon untuk bayar tepat waktu<br>4. Voucher ini hanya berlaku untuk pembayaran cicilan pertama<br>"
old_user_give_out_coupons_rule:"1. Kupon ini hanya berlaku untuk pemotongan 2% dari biaya layanan dan bukan untuk bunga pinjaman<br>2. Hanya berlaku untuk produk 90 hari<br>3. Hanya berlaku selama 90 hari dan kadaluarsa jika lebih dari 90 hari. Mohon untuk bayar tepat waktu<br>4. Voucher ini hanya berlaku untuk pembayaran cicilan pertama<br>"
old_user_push_coupons_rule:"1. Kupon ini hanya berlaku untuk pemotongan 2% dari biaya layanan dan bukan untuk bunga pinjaman<br>2. Hanya berlaku untuk produk 60, 90 dan 120 hari<br>3. Hanya berlaku selama 90 hari dan kadaluarsa jika lebih dari 90 hari. Mohon untuk bayar tepat waktu<br>4. Voucher ini hanya berlaku untuk pembayaran cicilan pertama<br>"

"1. Voucher 100 ribu rupiah ini dapat digunakan untuk diskon biaya pelayanan dan dapat digunakan dengan meminjam pinjaman<br>2. Hanya berlaku untuk produk 60, 90 dan 120 hari<br>3. Hanya berlaku selama 90 hari dan kadaluarsa jika lebih dari 90 hari. Mohon untuk bayar tepat waktu<br>4. Voucher ini hanya berlaku untuk pembayaran cicilan pertama<br>"
"1. voucher 100 ribu rupiah ini dapat digunakan untuk diskon biaya pelayanan dan dapat digunakan dengan meminjam pinjaman<br>2. Hanya berlaku untuk produk 60, 90 dan 120 hari<br>3. Hanya berlaku selama 90 hari dan kadaluarsa jika lebih dari 90 hari. Mohon untuk bayar tepat waktu    4. Voucher ini hanya berlaku untuk pembayaran cicilan pertama"


"1. Kupon ini hanya berlaku untuk pemotongan 2% dari biaya layanan dan bukan untuk bunga pinjaman2. Hanya berlaku untuk produk 90 hari3. Hanya berlaku selama 90 hari dan kadaluarsa jika lebih dari 90 hari. Mohon untuk bayar tepat waktu4. Voucher ini hanya berlaku untuk pembayaran cicilan pertama"
"1. Kupon ini hanya berlaku untuk pemotongan 2% dari biaya layanan dan bukan untuk bunga pinjaman2. Hanya berlaku untuk produk 90 hari3. Hanya berlaku selama 90 hari dan kadaluarsa jika lebih dari 90 hari. Mohon untuk bayar tepat waktu4. Voucher ini hanya berlaku untuk pembayaran cicilan pertama"
"1. Voucher 100 ribu rupiah ini dapat digunakan untuk diskon biaya pelayanan dan dapat digunakan dengan meminjam pinjaman<br>2. Hanya berlaku untuk produk 60, 90 dan 120 hari<br>3. Hanya berlaku selama 90 hari dan kadaluarsa jika lebih dari 90 hari. Mohon untuk bayar tepat waktu<br>4. Voucher ini hanya berlaku untuk pembayaran cicilan pertama<br>"


"1. Kupon ini hanya berlaku untuk pemotongan 2% dari biaya layanan dan bukan untuk bunga pinjaman<br>2. Hanya berlaku untuk produk 60, 90 dan 120 hari<br>3. Hanya berlaku selama 90 hari dan kadaluarsa jika lebih dari 90 hari. Mohon untuk bayar tepat waktu<br>4. Voucher ini hanya berlaku untuk pembayaran cicilan pertama<br>"


	host := viper.GetString("service.openapi.host")
  	iconMap := viper.GetStringMapString("monthly_pay_icon.app_commerce_icon")




大转盘活动

CREATE TABLE `user_withdrawal_record_form` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `uid` bigint(20) NOT NULL DEFAULT 0 COMMENT '用户id',
  `number_of_lucky_draws` int(11) NOT NULL DEFAULT 3 COMMENT '用户抽奖次数',
  `withdrawal` bigint(20) NOT NULL DEFAULT 0 COMMENT '抽奖提额',
  `status` int(11) NOT NULL DEFAULT 2 COMMENT '用户是否中过奖 1是 2否',
  `is_roll` int(11) NOT NULL DEFAULT 2 COMMENT '提额是否滚动过 1是 2否',
  `created` bigint(20) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated` bigint(20) NOT NULL DEFAULT '0' COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `idx_uid` (`uid`) USING BTREE,
  KEY `idx_created` (`created`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='用户提额记录表';






---老用户无uid push推送

select * from (select `u_id` as uid, device_id from (select uid,tmp.repayment_real_time from (select uid as u,repayment_real_time,count(id) as num from user_loan where  status = 4 and repayment_real_time>= UNIX_TIMESTAMP(DATE_ADD(NOW(), INTERVAL -48 HOUR)) and product_id =1 GROUP BY uid having num >= 1 ) as tmp,user_loan as a where a.uid =tmp.u and a.status not in (1,3) and a.status > 0 group by a.uid ) as temp,device as d where d.u_id = temp.uid group by d.u_id) as tmps

SELECT * FROM `google_event` WHERE `event_name` = 'new_user_limit' and created between 1681920000 and 1682524800



/* 16:49:36 pro-uangme_data uangme_data */ SELECT event_value,count('id') as num  FROM `google_event` WHERE `event_name` = 'new_user_limit' and created between 1681920000 and 1682524800 group by `event_value`

/* 16:49:36 pro-uangme_data uangme_data */ SELECT *  FROM `google_event` as ge left join user_info as ui on ge.uid = ui.`id` WHERE ge.event_name = 'new_user_limit' and ge.created between 1681920000 and 1682524800 and ui.promote_platform = 2

/* 16:49:36 pro-uangme_data uangme_data */ SELECT event_value,count('id') as num   FROM `google_event` WHERE `event_name` = 'new_user_limit' and created between 1682676469 and 1682762869 group by `event_value`

/* 16:49:36 pro-uangme_data uangme_data */ SELECT *   FROM `google_event` WHERE `event_name` = 'new_user_limit' and created between 1682676469 and 1682762869 


{
    "app_user_lev":1,
    "antif_006_pi_contacted_3d_cnt":0,
    "antif_208_gps_same_cnt_maxscore":0,
    "anti_053_contactpi_same_180d_overdue_30d_cnt":0,
    "exter_career_factory_worker":2,
    "exter_compete_level":1,
    "anti_071_gps_same_3d_cnt":0,
    "antif_012_pi_contacted_180d_overdue_1d_cnt":0,
    "exter_not_ojk_num":2,
    "antif_015_pi_contacted_90d_overdue_1d_cnt":0,
    "anti_072_gps_same_his_overdue_30d_cnt":0,
    "user_first_ab_score":51,
    "anti_077_gps_same_90d_overdue_7d_cnt":0,
    "exter_user_level":3,
    "antif_013_pi_contacted_90d_overdue_30d_cnt":0,
    "anti_053_contactpi_same_90d_cnt_score":0,
    "anti_053_contactpi_same_30d_overdue_1d_cnt":0,
    "exter_jia_have_fdc":1,
    "anti_073_gps_same_180d_overdue_30d_cnt":0,
    "exter_jia_app_black_loan_num":2,
    "exter_user_risk_lev_v2":3,
    "anti_053_contactpi_same_7d_cnt_score":0,
    "anti_052_contactpi_registered_his_overdue_1d_cnt":0,
    "anti_081_gps_same_90d_overdue_1d_cnt":0,
    "antif_301_anti_risk_finall_score_v1":0,
    "anti_052_contactpi_registered_180d_cnt":0,
    "exter_jia_fdc_score":528.7453944817731,
    "antif_017_pi_contacted_30d_overdue_1d_cnt_score":0,
    "antif_010_pi_contacted_180d_overdue_30d_cnt":0,
    "exter_tx_overdue":-999,
    "anti_052_contactpi_registered_90d_overdue_30d_cnt":0,
    "anti_052_contactpi_registered_30d_cnt":0,
    "anti_053_contactpi_same_180d_overdue_7d_cnt":0,
    "anti_066_gps_same_his_cnt_score":0,
    "antif_016_pi_contacted_30d_overdue_7d_cnt":0,
    "antif_005_pi_contacted_7d_cnt_score":0,
    "anti_052_contactpi_registered_7d_cnt":0,
    "exter_tx_refused_advance_score_rule":-999,
    "exter_advance_custom_score_rule":-999,
    "anti_052_contactpi_registered_180d_overdue_7d_cnt":0,
    "exter_hist_specified_amount":-1,
    "fdc_last_1month_old_amt":2370.0000000000005,
    "exter_channel_flag":3,
    "anti_052_contactpi_registered_90d_cnt":0,
    "exter_jia_loan_app_num":5,
    "exter_channel_channel":100054,
    "anti_078_gps_same_30d_overdue_7d_cnt":0,
    "exter_watsApp_result":-999,
    "anti_080_gps_same_180d_overdue_1d_cnt":0,
    "antif_201_pi_contacted_overdue_cnt_maxscore":0,
    "hit_strategy_id":-1,
    "exter_jia_gps_30d_same_num":1,
    "anti_069_gps_same_30d_cnt":0,
    "exter_channel_promote_agency":56,
    "anti_052_contactpi_registered_30d_overdue_7d_cnt":0,
    "anti_084_email_same_his_cnt_score":0,
    "exter_jia_fdc_combine_score":532.0896669398212,
    "anti_052_contactpi_registered_90d_cnt_score":0,
    "exter_jia_wifi_7d_reg_user_num":1,
    "anti_053_contactpi_same_his_cnt":0,
    "exter_jia_multi_score":-999,
    "anti_052_contactpi_registered_30d_overdue_1d_cnt_score":0,
    "exter_rupiah_plus_is_d7":-999,
    "anti_079_gps_same_his_overdue_1d_cnt":0,
    "anti_070_gps_same_7d_cnt":0,
    "anti_057_dev_same_30d_cnt_score":0,
    "exter_user_risk_segment_v2":42,
    "anti_055_dev_same_180d_cnt":0,
    "exter_use_custom_change_score":-999,
    "anti_053_contactpi_same_30d_overdue_7d_cnt":0,
    "anti_052_contactpi_registered_his_overdue_7d_cnt":0,
    "anti_067_gps_same_180d_cnt_score":0,
    "antif_009_pi_contacted_his_overdue_1d_cnt_score":0,
    "anti_052_contactpi_registered_90d_overdue_1d_cnt":0,
    "antif_007_pi_contacted_his_overdue_30d_cnt":0,
    "exter_jia_advance_apply_count":4,
    "anti_053_contactpi_same_30d_cnt":0,
    "exter_expert_combine_score":554,
    "exter_easycash_flag":1,
    "anti_067_gps_same_180d_cnt":0,
    "anti_052_contactpi_registered_90d_overdue_7d_cnt":0,
    "exter_jia_phone_relation_30d_contact_num":0,
    "exter_tx_loan_refused":-999,
    "antif_004_pi_contacted_30d_cnt_score":0,
    "antif_001_pi_contacted_his_cnt":0,
    "exter_jia_app_score":-999,
    "exter_id_card_check_is_vaild":2,
    "antif_011_pi_contacted_180d_overdue_7d_cnt":0,
    "anti_053_contactpi_same_7d_cnt":0,
    "anti_053_contactpi_same_30d_cnt_score":0,
    "anti_058_dev_same_7d_cnt":0,
    "anti_052_contactpi_registered_30d_cnt_score":0,
    "anti_053_contactpi_same_90d_cnt":0,
    "anti_068_gps_same_90d_cnt":0,
    "anti_053_contactpi_same_his_overdue_1d_cnt":0,
    "antif_017_pi_contacted_30d_overdue_1d_cnt":0,
    "exter_jia_contact_30d_same_num":1,
    "anti_052_contactpi_registered_30d_overdue_1d_cnt":0,
    "exter_legal_ojk_num2":3,
    "exter_app_level":2,
    "exter_old_model_strategy_change_score":-999,
    "exter_hist_is_refuse":"NO",
    "exter_jia_wifi_3d_reg_user_num":1,
    "antif_006_pi_contacted_3d_cnt_score":0,
    "exter_whatapp_regist_phone_is_same":1,
    "exter_full_address_w_len":20,
    "exter_ui_create_hour":"06",
    "exter_use_custom_change":-999,
    "uid":11523544,
    "anti_053_contactpi_same_his_cnt_score":0,
    "anti_083_gps_same_7d_overdue_1d_cnt":0,
    "anti_052_contactpi_registered_his_cnt":0,
    "anti_053_contactpi_same_3d_cnt":0,
    "exter_top_notlegal_ojk_num":0,
    "exter_advance_ai_multi_score":-999,
    "antif_005_pi_contacted_7d_cnt":0,
    "antif_008_pi_contacted_his_overdue_7d_cnt":0,
    "anti_053_contactpi_same_180d_overdue_1d_cnt":0,
    "anti_054_dev_same_his_cnt":0,
    "antif_205_contactpi_registered_overdue_cnt_maxscore":0,
    "anti_052_contactpi_registered_180d_overdue_1d_cnt":0,
    "antif_001_pi_contacted_his_cnt_score":0,
    "exter_tx_overdays_advance_score_rule":-999,
    "anti_074_gps_same_90d_overdue_30d_cnt":0,
    "exter_strategy_flag":1,
    "exter_user_amount_rate":44,
    "anti_059_dev_same_3d_cnt":0,
    "anti_075_gps_same_his_overdue_7d_cnt":0,
    "anti_057_dev_same_30d_cnt":0,
    "fdc_multi_period_last_6m_loan_cnt":7,
    "exter_jia_contact_15d_same_num":1,
    "anti_053_contactpi_same_90d_overdue_7d_cnt":0,
    "exter_jia_phone_is_same_cotact_phone":0,
    "exter_rupiah_plus_max_overdays":-999,
    "anti_053_contactpi_same_90d_overdue_1d_cnt":0,
    "exter_rule_implement":0,
    "fdc_multi_period_last_6m_loan_max_overdue_days":0,
    "anti_052_contactpi_registered_7d_cnt_score":0,
    "exter_advance_ai_fraud_score":-999,
    "exter_channel_name_flag":6,
    "anti_052_contactpi_registered_180d_overdue_30d_cnt":0,
    "exter_amount_limit_flag":-1,
    "anti_053_contactpi_same_his_overdue_7d_cnt":0,
    "exter_all_ojk_num":3,
    "anti_053_contactpi_same_180d_cnt":0,
    "exter_jia_name_birth_same_num":1,
    "anti_066_gps_same_his_cnt":0,
    "exter_ios_base_info_level":5,
    "exter_fdc_user_trans_flag":1,
    "exter_user_risk_lev":2,
    "antif_206_contactpi_same_overdue_cnt_maxscore":0,
    "exter_jia_app_black_dc_num":0,
    "exter_rupiah_plus_is_d14":-999,
    "antif_207_dev_same_cnt_maxscore":0,
    "anti_052_contactpi_registered_his_overdue_30d_cnt":0,
    "antif_004_pi_contacted_30d_cnt":0,
    "antif_009_pi_contacted_his_overdue_1d_cnt":0,
    "exter_fdc_user_flag":-3,
    "anti_052_contactpi_registered_his_cnt_score":0,
    "anti_053_contactpi_same_90d_overdue_30d_cnt":0,
    "exter_wa_result_advance_score_rule":-999,
    "anti_076_gps_same_180d_overdue_7d_cnt":0,
    "exter_jia_no_fdc_combine_score":532.2658688525473,
    "anti_056_dev_same_90d_cnt":0,
    "anti_052_contactpi_registered_3d_cnt":0,
    "exter_ad_7d_60d_querypct":0,
    "exter_campaign_flag":-1,
    "exter_product_code":108,
    "antif_003_pi_contacted_90d_cnt":0,
    "anti_053_contactpi_same_his_overdue_30d_cnt":0,
    "anti_054_dev_same_his_cnt_score":0,
    "anti_082_gps_same_30d_overdue_1d_cnt":0,
    "anti_084_email_same_his_cnt":0,
    "exter_channel_promote_platform":58,
    "antif_002_pi_contacted_180d_cnt":0,
    "exter_jia_device_same_num":1,
    "exter_strategy_change_score":92.88227,
    "antif_014_pi_contacted_90d_overdue_7d_cnt":0,
    "exter_jia_wifi_15d_reg_user_num":1
}


----排查问题sql

/* 19:00:57 push库 cmpush_in */ SELECT count(id) as num  FROM `hokuto_msg_topic` WHERE `description` = 'not find token' and created between 1683457308 and 1683543708 ORDER BY `id` DESC LIMIT 1;

select FROM_UNIXTIME(created,'%Y-%m-%d') as dt ,count(id) as num from hokuto_msg_topic  WHERE `description` = 'not find token' and created > 1682870448 group by dt 

select FROM_UNIXTIME(created,'%Y-%m-%d') as dt ,count(id) as num from hokuto_msg_topic  WHERE `description` = 'not find token' and created > 1683648000 group by dt 



select push_id,count(id) as num from hokuto_msg_topic  WHERE `description` = 'not find token' and created between 1683561648 and 1683648048 group by push_id order by num desc

select FROM_UNIXTIME(created,'%Y-%m-%d') as dt ,count(id) as num from hokuto_uangme_token  WHERE created > 1681920000 group by dt 


==========================================================================================================================================================================================
---- 新版google打点
/* 11:01:23 pro-uangme_loan uangme_loan */ ALTER TABLE `user_rc_info` ADD `user_risk_lev_v2` INT(11)  NOT NULL  DEFAULT '0'  COMMENT '用户风控等级V2'  AFTER `user_risk_lev`;

CREATE TABLE `user_new_reg` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `uid` bigint(20) NOT NULL DEFAULT 0 COMMENT '用户id',
  `app_version` char(40) NOT NULL DEFAULT 0 COMMENT '用户注册app版本',
  `status` int(11) NOT NULL DEFAULT 1 COMMENT '用户是否正常 1是 2否',
  `created` bigint(20) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated` bigint(20) NOT NULL DEFAULT '0' COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `idx_uid` (`uid`) USING BTREE,
  KEY `idx_created` (`created`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='用户注册信息版本记录表';




hive_sync_uangme_loan


CREATE TABLE `hive_sync_uangme_loan` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `uid` bigint(20) NOT NULL DEFAULT 0 COMMENT '用户ID',
  `last_apply_date` bigint(20) NOT NULL DEFAULT 0 COMMENT '最后一次进件日期',
  `amount` int(11) NOT NULL DEFAULT 0 COMMENT '最后一次授信通过后的额度',
  `is_pass` char(20) NOT NULL DEFAULT 1 COMMENT '风控通过:YES',
  `huoe_to_today` bigint(20) NOT NULL DEFAULT 0 COMMENT '当天和最后一次进件日期的比较日期',
  `exter_user_risk_lev_v2` int(11) NOT NULL DEFAULT 0 COMMENT '用户风控V2等级',
  `ina_tx_time` bigint(20) NOT NULL DEFAULT 0 COMMENT '提现时间 减一个小时',
  `ext` char(255) NOT NULL DEFAULT '' COMMENT '冗余字段',
  `real_repayment_dt` bigint(20) NOT NULL DEFAULT 0 COMMENT '用户还款时间 减 一个小时',
  `repay_to_tx_diff` bigint(20) NOT NULL DEFAULT 0 COMMENT '用户最近一次提现距离还款时间',
  `exter_dz_model_long_v2_score` int(11) NOT NULL DEFAULT 0 COMMENT '长账龄用户评分',
  `dz_level_short_v2` int(11) NOT NULL DEFAULT 0 COMMENT '短账龄用户评级',
  `exter_dz_model_short_v2_score` int(11) NOT NULL DEFAULT 0 COMMENT 'exter_dz_model_short_v2_score 未知含义',
  `exter_first_trans_to_apply_days` int(11) NOT NULL DEFAULT 01 COMMENT '用户账龄字段',
  `total_amount` int(11) NOT NULL DEFAULT 0 COMMENT '用户当前额度',
  `pass_cnt` int(11) NOT NULL DEFAULT 0 COMMENT '当前获额状态',
  `zd_cnt` int(11) NOT NULL DEFAULT 0 COMMENT '当前是否在贷',
  `is_overdueing_flag` int(11) NOT NULL DEFAULT 0 COMMENT '当前是否逾期中',
  `created` bigint(20) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated` bigint(20) NOT NULL DEFAULT '0' COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `idx_uid` (`uid`) USING BTREE,
  KEY `idx_created` (`created`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='风控hive同步数据表';




[[job_list]]
    time = "0 0 6-23 * * *"  # 每小时
    title = "uangme_data业务系统告警:第三方izidata请求失败率过高1"
    job_name = "UnifiedIndexMonitor"
    status = false
    [job_list.params]
        type = 0
        command = "SELECT round((S1.num/S2.total)*100,0) AS result FROM ( SELECT count('id') as num FROM `third_req_res_record` WHERE `type` = '2' and `code` = 'FAIL' ) S1,( SELECT count('id') as total FROM `third_req_res_record` WHERE `type` = '2') S2;"
        
        threshold = 5
        duration = 3600
        database = "uangme_data"
        [[job_list.params.column]]
            msg = ""
            min = 10
            max = 90000
            key = "result"
            sms_to = ["13321188634"]
            dtalk_to = ["13321188634"]
            sub_sql = "liangpengzhan@superatomfin.com"



[[job_list]]
    time = "0 */1 * * * *"  # 秒 分 时 日 月 周
    title = "业务系统指标告警:cm_dai_admin[强规则拒绝率过高]1111"
    job_name = "UnifiedIndexMonitor"
    status = true      # 是否启用
    [job_list.params]
        type = 0       # 类型, 0表示command为sql， 1为自定义handler(需要代码内添加函数)
        command = "select count(distinct user_id) as data from audit_record where is_effective =1 and strong_rule not in (0,6) and created >= UNIX_TIMESTAMP(DATE_SUB(now(),INTERVAL 60 MINUTE)) group by strong_rule having data > 50 limit 1;"
        database = "uangme"
        threshold = 1 # 告警阈值
        duration = 10  # 持续时间, 单位分钟
        [[job_list.params.column]]
            msg = ""     # 告警文案
            min = 0       # 浮点数最小值
            max = 50       # 浮点数最大值
            key = "data"    # 取值的key, 对应sql中输出的data字段或handler输出的data字段
            mail_to = ["guoshaohua@superatomfin.com"]    # 邮件接件人
            # sms_to = ["13520210930","18600013217"]     # 短信接收人
            dtalk_to = ["13520210930","18600013217"]


----排查sql  2023年06月06日 星期二 开始

-- 正式的
1.是否再贷

select * from user_loan where uid = 1111 and status in (0,1,3)

2.获取用户账期
select get_json_object(data,"$.exter_first_trans_to_apply_days") as exter_first_trans_to_apply_days from rc_outdata_repayment_result where uid = 1111

select * from task_log where created >=1685950800 and task_id = 1156

select * from task_log where created >=1686019200 and task_id = 1155

select * from task_log where uid = 1532971 and task_id = 1155 AND created >=1686019200

select * from (select uid%100 as a,created,task_id from task_log as t where created >=1686019200 and task_id = 1155 ) as i where i.a <= 5

	select FROM_UNIXTIME(created,'%Y-%m-%d') as dt ,count(id) as num from hokuto_msg_topic  WHERE `description` = 'not find token' and created > 1685462400 group by dt 


SELECT round((S1.num/S2.total)*100,0) AS data FROM ( SELECT count('id') as num FROM `third_req_res_record` WHERE `type` = '2' and `code` = 'FAIL' ) S1,( SELECT count('id') as total FROM `third_req_res_record` WHERE `type` = '2') S2;

select FROM_UNIXTIME(created,'%Y-%m-%d') as dt ,count(id) as num from hokuto_msg_topic  WHERE `description` = 'not find token' and created > 1682870448 group by dt 


select count(id) as num from user_rc_info where created BETWEEN 1685376000 and 1685462402 and user_risk_lev_v2 BETWEEN 1 and 3


select count(id) as num from app_flyer_event where created BETWEEN 1685376000 and 1685462400 and event_name = 'new_user_level_V2_1_3' and media_source = 'googleadwords_int'

select count(id) as num from app_flyer_event where created BETWEEN 1685376000 and 1685462400 and event_name = 'new_user_level_V2_1_3' and media_source = 'Facebook Ads'

select count(id) as num from app_flyer_event where created BETWEEN 1685376000 and 1685462400 and event_name = 'new_user_level_V2_1_3' and media_source = 'bytedanceglobal_int'


select DISTINCT(uid) from user_rc_info where created BETWEEN 1685376000 and 1685462402 and user_risk_lev_v2 BETWEEN 1 and 3 and ext = 'true' order by uid desc 
	select FROM_UNIXTIME(created,'%Y-%m-%d') as dt ,count(id) as num from hokuto_msg_topic  WHERE `description` = 'not find token' and created > 1685462400 group by dt 

	select FROM_UNIXTIME(created,'%Y-%m-%d') as dt ,count(id) as num from hokuto_msg_topic  WHERE `description` = 'not find token' and created > 1685462400 group by dt 

/* 10:49:24 pro-uangme_loan uangme_loan */ SELECT count('id') as num FROM `user_withdrawal_record_form` WHERE `status` = '1' LIMIT 0,1000;


select * from user_info a LEFT JOIN user_loan b on a.id=b.uid where a.last_submit_time BETWEEN UNIX_TIMESTAMP(DATE_SUB(CURRENT_DATE,INTERVAL 5 DAY)) AND UNIX_TIMESTAMP(DATE_SUB(CURRENT_DATE,INTERVAL 4 DAY))
and a.verify_status=2 and b.loan_id is NULL


select count(case when strong_rule not in (0,6) then user_id end)/count(distinct user_id) as data from audit_record where is_effective =1  and created >= UNIX_TIMESTAMP(DATE_SUB(now(),INTERVAL 60 MINUTE));

select count(distinct user_id) as data from audit_record where is_effective =1 and strong_rule not in (0,6) and created >= UNIX_TIMESTAMP(DATE_SUB(now(),INTERVAL 60 MINUTE)) group by strong_rule having data > 50 limit 1;

SELECT round((S1.num/S2.total)*100,0) AS result FROM ( SELECT count('id') as num FROM `third_req_res_record` WHERE `type` = '2' and `code` = 'FAIL' ) S1,( SELECT count('id') as total FROM `third_req_res_record` WHERE `type` = '2') S2;

select * from user_info a LEFT JOIN user_loan b on a.id=b.uid where a.last_submit_time BETWEEN UNIX_TIMESTAMP(DATE_SUB(CURRENT_DATE,INTERVAL 5 DAY)) AND UNIX_TIMESTAMP(DATE_SUB(CURRENT_DATE,INTERVAL 4 DAY))
and a.verify_status=2 and b.loan_id is NULL

SELECT round((S1.num/S2.total)*100,0) AS data FROM ( SELECT count('id') as num FROM `third_req_res_record` WHERE `type` = '2' and `code` = 'FAIL' ) S1,( SELECT count('id') as total FROM `third_req_res_record` WHERE `type` = '2') S2;

select a.id as uid,c.total_amount as amount,d.product_id as pid,date(from_unixtime(a.last_submit_time-3600)) as loan_date  from user_info a LEFT JOIN user_loan b on a.id=b.uid LEFT JOIN product_account c on a.id=c.uid LEFT JOIN user_product d on a.id=d.uid where a.last_submit_time BETWEEN UNIX_TIMESTAMP(DATE_SUB(CURRENT_DATE,INTERVAL 5 DAY)) AND UNIX_TIMESTAMP(DATE_SUB(CURRENT_DATE,INTERVAL 4 DAY))
and a.verify_status=2 and b.loan_id is NULL and c.account_type=2 and d.is_pass='YES' group by a.id


select DISTINCT a.uid,c.product_id as pid,d.total_amount as amount,date(from_unixtime(rtime-3600)) as loan_date from
(select uid,MAX(real_repayment_time) as rtime from user_loan_plan where real_repayment_time BETWEEN UNIX_TIMESTAMP(DATE_SUB(CURRENT_DATE,INTERVAL 7 DAY)) AND UNIX_TIMESTAMP(DATE_SUB(CURRENT_DATE,INTERVAL 6 DAY))
and `status`=2 GROUP BY uid
) a LEFT JOIN user_loan_plan b on a.uid=b.uid and b.ctime>a.rtime LEFT JOIN user_product c on a.uid=c.uid LEFT JOIN product_account d on a.uid=d.uid where c.is_pass='YES' and c.product_id in (108,119,116,120,121) and b.id is NULL and d.account_type=2;

SELECT round((S1.num/S2.total)*100,0) AS result FROM ( SELECT count('id') as num FROM `third_req_res_record` WHERE `type` = '2' and `code` = 'FAIL' ) S1,( SELECT count('id') as total FROM `third_req_res_record` WHERE `type` = '2') S2;


select uid from google_event where created BETWEEN 1685376000 and 1685462402 and event_name = 'new_user_level_V2_1_3' order by uid desc

select uid from google_event where uid =11724150 and event_name = 'new_user_level_V2_1_3'

select a.id as uid,c.total_amount as amount,d.product_id as pid from user_info a LEFT JOIN user_loan b on a.id=b.uid LEFT JOIN product_account c on a.id=c.uid LEFT JOIN user_product d on a.id=d.uid where a.last_submit_time BETWEEN UNIX_TIMESTAMP(DATE_SUB(CURRENT_DATE,INTERVAL 5 DAY)) AND UNIX_TIMESTAMP(DATE_SUB(CURRENT_DATE,INTERVAL 4 DAY))
and a.verify_status=2 and b.loan_id is NULL and c.account_type=2 and d.is_pass='YES' group by a.id

select * from user_loan_plan where uid = 11682741 and status in (0,1) and overdue_days>=1


-- 测试的

select * from amount_change where uid = 7942754 and created >= 1684060320 and method in ('recall')

select * from amount_change where uid in (7942754,7944576,7944567,7944212,7944134) and created >= 1684060320 and method in ('recall') group by uid

select * from user_loan_plan where uid in ( 7942754,7944576,7944567,7944212,7944134) and status in (0,1,3) group by uid

select id,uid,status from user_loan_plan where uid in ( 7942754,7944576,7944567,7944212,7944134) and status in (0,1,3) 

select * from user_loan_plan where uid in (7942754,7944576,7944567,7944212,7944134) and status in (0,1,3) 

select * from user_loan_plan where uid = 7942754 and status in (0,1,3)


select *,JSON_EXTRACT(data,"$.exter_dz_model_long_v2_score") as exter_dz_model_long_v2_score,JSON_EXTRACT(data,"$.exter_dz_model_short_v2_score") as exter_dz_model_short_v2_score from crm_feature.rc_outdata_model_data t where type = "exter_jia_dz_v2" and uid in(100596)

select description,substring_index(substring_index(description,'"result":',-1),',"',1) AS 'result:' from hokuto_msg_topic limit 1



INSERT INTO `market_strategy` (`strategy_id`, `strategy_name`, `project`, `filters`, `action`, `status`, `is_test`, `created`, `updated`)
VALUES
	(1, '注册成功', 'uangme', NULL, '[{\"label\":\"market_event执行器\",\"name\":\"market_event_action\"}]', 1, 0, 0, 0),
	(2, '进件完成', 'uangme', NULL, '[{\"label\":\"market_event执行器\",\"name\":\"market_event_action\"}]', 1, 0, 0, 0),
	(3, 'Cash loan授信成功', 'uangme', NULL, '[{\"label\":\"market_event执行器\",\"name\":\"market_event_action\"}]', 1, 0, 0, 0),
	(4, 'Cash loan放款成功', 'uangme', NULL, '[{\"label\":\"market_event执行器\",\"name\":\"market_event_action\"}]', 1, 0, 0, 0),
	(5, '用户等级（多打）', 'uangme', NULL, '[{\"label\":\"market_event执行器\",\"name\":\"market_event_action\"}]', 1, 0, 0, 0),
	(6, '优质交易事件', 'uangme', NULL, '[{\"label\":\"market_event执行器\",\"name\":\"market_event_action\"}]', 1, 0, 0, 0),
	(7, '贷前V2评级打点事件', 'uangme', NULL, '[{\"label\":\"market_event执行器\",\"name\":\"market_event_action\"}]', 1, 0, 0, 0),
	(8, 'LTV数值事件', 'uangme', NULL, '[{\"label\":\"market_event执行器\",\"name\":\"market_event_action\"}]', 1, 0, 0, 0),
	(9, '用户等级（单打）	', 'uangme', NULL, '[{\"label\":\"market_event执行器\",\"name\":\"market_event_action\"}]', 1, 0, 0, 0),
	(11, '提交借款', 'uangme', NULL, '[{\"label\":\"market_event执行器\",\"name\":\"market_event_action\"}]', 1, 0, 0, 0),
	(14, 'PayLater授信成功	', 'uangme', NULL, '[{\"label\":\"market_event执行器\",\"name\":\"market_event_action\"}]', 1, 0, 0, 0),
	(15, 'PayLater放款成功', 'uangme', NULL, '[{\"label\":\"market_event执行器\",\"name\":\"market_event_action\"}]', 1, 0, 0, 0),
	(16, '强规则通过', 'uangme', NULL, '[{\"label\":\"market_event执行器\",\"name\":\"market_event_action\"}]', 1, 0, 0, 0),
	(17, '高分', 'uangme', NULL, '[{\"label\":\"market_event执行器\",\"name\":\"market_event_action\"}]', 1, 0, 0, 0),
	(18, '中分', 'uangme', NULL, '[{\"label\":\"market_event执行器\",\"name\":\"market_event_action\"}]', 1, 0, 0, 0),
	(19, '低分', 'uangme', NULL, '[{\"label\":\"market_event执行器\",\"name\":\"market_event_action\"}]', 1, 0, 0, 0),
	(20, '强规则拒绝', 'uangme', NULL, '[{\"label\":\"market_event执行器\",\"name\":\"market_event_action\"}]', 1, 0, 0, 0),
	(21, '多次放款成功', 'uangme', NULL, '[{\"label\":\"market_event执行器\",\"name\":\"market_event_action\"}]', 1, 0, 0, 0);

  /* 15:22:34 dev opadmin */ UPDATE `market_event_rule` SET `template` = null WHERE `strategy_id` = '5' and `source` = 'google' and `event_id` between 5 and 11;




CREATE TABLE `not_active_user_chang_amount_record_form` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `uid` bigint(20) NOT NULL DEFAULT 0 COMMENT '用户ID',
  `user_amount` bigint(20) NOT NULL DEFAULT 0 COMMENT '用户额度',
  `scene` int(11) NOT NULL DEFAULT 0 COMMENT '场景 1是贷前 2贷中',
  `status` int(11) NOT NULL DEFAULT 0 COMMENT '状态 1正常 2异常',
  `is_gray` int(11) NOT NULL DEFAULT 0 COMMENT '是否灰度 1 是 2否',
  `created` bigint(20) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated` bigint(20) NOT NULL DEFAULT '0' COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `idx_uid` (`uid`) USING BTREE,
  KEY `idx_created` (`created`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='未激活用户提额记录表';



  -- `is_pass` char(20) NOT NULL DEFAULT 1 COMMENT '风控通过:YES',
  -- `huoe_to_today` bigint(20) NOT NULL DEFAULT 0 COMMENT '当天和最后一次进件日期的比较日期',
  -- `exter_user_risk_lev_v2` int(11) NOT NULL DEFAULT 0 COMMENT '用户风控V2等级',
  -- `ina_tx_time` bigint(20) NOT NULL DEFAULT 0 COMMENT '提现时间 减一个小时',
  -- `ext` char(255) NOT NULL DEFAULT '' COMMENT '冗余字段',
  -- `real_repayment_dt` bigint(20) NOT NULL DEFAULT 0 COMMENT '用户还款时间 减 一个小时',
  -- `repay_to_tx_diff` bigint(20) NOT NULL DEFAULT 0 COMMENT '用户最近一次提现距离还款时间',
  -- `exter_dz_model_long_v2_score` int(11) NOT NULL DEFAULT 0 COMMENT '长账龄用户评分',
  -- `dz_level_short_v2` int(11) NOT NULL DEFAULT 0 COMMENT '短账龄用户评级',
  -- `exter_dz_model_short_v2_score` int(11) NOT NULL DEFAULT 0 COMMENT 'exter_dz_model_short_v2_score 未知含义',
  -- `exter_first_trans_to_apply_days` int(11) NOT NULL DEFAULT 01 COMMENT '用户账龄字段',
  -- `total_amount` int(11) NOT NULL DEFAULT 0 COMMENT '用户当前额度',
  -- `pass_cnt` int(11) NOT NULL DEFAULT 0 COMMENT '当前获额状态',
  -- `zd_cnt` int(11) NOT NULL DEFAULT 0 COMMENT '当前是否在贷',
  -- `is_overdueing_flag` int(11) NOT NULL DEFAULT 0 COMMENT '当前是否逾期中',



        {
            "label": "用户风控评分过滤器",
            "name": "user_risk_score_filter",
            "params": {
                "score": {
                    "label": "最低分值",
                    "field_type": "input",
                    "type": "int",
                    "value": "620"
                }
            }
        },
        {
            "label": "用户账龄过滤器",
            "name": "user_account_age_filter",
            "params": {
                "condition": {
                    "label": "条件",
                    "field_type": "input",
                    "type": "",
                    "value": ">="
                },
                "day": {
                    "label": "账龄(天)",
                    "field_type": "input",
                    "type": "int",
                    "value": "120"
                }
            }
        },


        "level": {
                    "label": "用户当前等级",
                    "field_type": "tag_input",
                    "type": "",
                    "value": "0"
                },

/* 14:26:59 dev opadmin */ ALTER TABLE `task_log` CHANGE `event_result` `event_result` VARCHAR(2000)  CHARACTER SET utf8  COLLATE utf8_general_ci  NOT NULL  DEFAULT ''  COMMENT '执行事件的结果json';
