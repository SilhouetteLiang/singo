
1010    300000 
1011    330000
1012    370000
1013    410000
1014    450000
1015    480000
1016    520000


1350
1351
1352
1353
1354
1355
1356


Plafon Pinjaman Telah Naik!
Content: 
Karena riwayatmu baik, kami sudah menaikkan plafonmu dari Rp{{$ds.amount}}menjadi Rp{{$ds.new_amount}}. Ayo buka UangMe untuk bisa cairkan limit barumu!


Pinjam Lagi Dengan Limit Lebih Besar!
Kuota Pinjamanmu sudah diupgrade loh dari Rp{{$ds.amount}} menjadi Rp{{$ds.new_amount}}! Ayo cairkan kembali di UangMe dan manfaatkan semaksimal mungkin!








INSERT INTO `market_event_field` (`source`, `name`, `value`, `target_name`, `target_value`, `created`, `updated`)
VALUES
	( 'google', 'ltv', '280', 'user_rc_lev_3', '1', 1690771655, 1690771655),
	( 'google', 'ltv', '250', 'user_rc_lev_3', '2', 1690771655, 1690771655),
	( 'google', 'ltv', '60', 'user_rc_lev_3', '3', 1690771655, 1690771655),
	( 'google', 'ltv', '34', 'user_rc_lev_3', '4', 1690771655, 1690771655),
	( 'google', 'ltv', '8', 'user_rc_lev_3', '5', 1690771655, 1690771655),
	( 'google', 'ltv', '0', 'user_rc_lev_3', '6', 1690771655, 1690771655),
	( 'google', 'ltv', '0', 'user_rc_lev_3', '7', 1690771655, 1690771655);

INSERT INTO `market_event_rule` (`strategy_id`, `action`, `filters`, `template`, `event_name`, `event_value`, `source`, `status`, `ext`, `created`, `updated`)
VALUES
	(4, 'market_event_action', '[{\"name\":\"user_rc_lev_3\",\"label\":\"用户等级过滤器\",\"params\":{\"user_rc_lev_3\":[1,2,3,4,5,6,7]}},{\"name\":\"promote_platform\",\"label\":\"渠道过滤器\",\"params\":{\"promote_platform\":[2]}}]', '{\"value\": \"{{$ds.market_event_field.ltv.user_rc_lev_3}}\"}', 'new_user_limit', '1009', 'google', 1, NULL, 0, 0);



1005老客营销普调PUSH 2250000≤限额<3000000 2023年07月29日 110000
1006老客营销普调PUSH 3000000≤限额<3750000 2023年07月29日 150000
1007老客营销普调PUSH 3750000≤限额<4500000 2023年07月29日 180000
1008老客营销普调PUSH 4500000≤限额<5250000 2023年07月29日 220000
1009老客营销普调PUSH 5250000≤限额<6000000 2023年07月29日 260000

1005老客营销普调PUSH 2250000≤限额<3000000 2023年07月29日
1006老客营销普调PUSH 3000000≤限额<3750000 2023年07月29日
1007老客营销普调PUSH 3750000≤限额<4500000 2023年07月29日
1008老客营销普调PUSH 4500000≤限额<5250000 2023年07月29日
1009老客营销普调PUSH 5250000≤限额<6000000 2023年07月29日

提额券ID    提额额度
1005        110000
1006        150000
1007        180000
1008        220000
1009        260000

1334-
1335
1336
1337
1338

1329-82319  28号
1330-82318  28号
1331-82317  28号
1332-82321  28号
1333-82320  28号

1329-83371  29号
1330-83372  29号
1331-83369  29号
1332-83370  29号
1333-83368  29号






1010老客营销普调PUSH 6000000≤限额<6750000   
1011老客营销普调PUSH 6750000≤限额<7500000
1012老客营销普调PUSH 7500000≤限额<8250000
1013老客营销普调PUSH 8250000≤限额<9000000
1014老客营销普调PUSH 9000000≤限额<9750000
1015老客营销普调PUSH 9750000≤限额<10500000
1016老客营销普调PUSH 6750000≤限额<7500000


SELECT DISTINCT(ual.uid) AS uid,
        pa.total_amount AS amount,
        pa.total_amount+260000 AS new_amount
FROM user_active_log AS ual
LEFT JOIN product_account pa
    ON ual.uid = pa.uid
LEFT JOIN user_rc_info uri
    ON ual.uid = uri.uid
LEFT JOIN user_loan ul
    ON ual.uid = ul.uid
WHERE ual.created
    BETWEEN UNIX_TIMESTAMP(DATE_SUB(CURRENT_DATE,INTERVAL 30 DAY))
        AND UNIX_TIMESTAMP()
        AND pa.use_amount/pa.total_amount >= 0.5
        AND uri.is_no_cycle = 0
        AND ul.status IN (3,4)
        AND pa.total_amount >= 5250000
        AND pa.total_amount <  6000000
ORDER BY  ul.uid 

[UangMe] Kuota Pinjamanmu sudah diupgrade loh dari Rp{{$ds.amount}} menjadi Rp ({{$ds.amount}}+110000)! Ayo cairkan kembali di UangMe dan manfaatkan semaksimal mungkin!


select count(id) as num from (
select id,message_id,state,substring_index(substring_index(description,'"result":',-1),',"',1) AS 'result' from hokuto_msg_topic where created between 1680278400 and 1682870400
) as t where t.result = false

2023年07月26日 星期三


select * from task_log where task_id in (1172,1173,1174,1199) and created=0 and status = 1 and event_type = 2


SELECT count(task_id) AS data FROM(SELECT count(id) AS num, event_type , task_id FROM task_log WHERE task_id in(1166,1185,1167,1172,1178,1173,1199,1179,1172,1175,1173,1176,1199,1200,1174,1177,1178,1179,1180,1181,1182,1183) AND status = 1 AND event_type=2 AND created >= UNIX_TIMESTAMP(DATE_SUB(CURRENT_DATE,INTERVAL -1 DAY))-16*3600 GROUP BY event_type, task_id) AS a

SELECT count(task_id) AS data FROM(SELECT count(id) AS num, event_type , task_id FROM task_log WHERE task_id in(1166, 1167, 1168, 1169, 1170, 1171, 1172, 1173, 1174, 1175, 1176, 1177, 1199, 1200, 1178, 1179, 1180, 1181, 1182, 1183) AND status = 1 AND event_type=2 AND created >= UNIX_TIMESTAMP(DATE_SUB(CURRENT_DATE,INTERVAL -1 DAY))-16*3600 GROUP BY event_type, task_id) AS a
  

update task_log set created = 1687742068 where task_id in (1172,1173,1174,1199) and created=0 and status = 1 and event_type = 2

UPDATE `task_log` SET `created` = '1687742068' WHERE `task_id` in (1172,1173,1174,1199) and created=0 and status = 1 and event_type = 2;

select * from task_log WHERE `task_id` in (1172,1173,1174,1199) and created=0 and status = 1 and event_type = 2;

UPDATE `task_log` SET `created` = '1687742068' WHERE `task_id` in (1172,1173,1174,1199,1175,1176,1177,1200,1178,1179,1180,1181,1182,1183) and created=0 and status = 1 and event_type = 2;

select count(id) as num,task_id from task_log  WHERE `task_id` in (1172,1173,1174,1199,1175,1176,1177,1200,1178,1179,1180,1181,1182,1183) and created>=1687742067 and status = 1 and event_type = 2  group by task_id;

select count(id) as num,task_id from task_log  WHERE `task_id` in (1172,1173,1174,1199,1175,1176,1177,1200,1178,1179,1180,1181,1182,1183) and created=0 and status = 1 and event_type = 1  group by task_id;



select * from task_log where event_name = 'strategy_account_amount_action' and  limit  100

select * from task_log  WHERE `task_id` in (1172,1173,1174,1199,1175,1176,1177,1200,1178,1179,1180,1181,1182,1183) and created=0 ;

select task_id,count(id) as num from task_log where created=0 group by task_id

select 
count(id) as data,
event_type ,
task_id
from task_log 
where task_id in (1172,1175,1173,1176,1199,1200,1174,1177,1178,1179,1180,1181,1182,1183) 
and status = 1  
and event_type=2
and created >= UNIX_TIMESTAMP(DATE_SUB(CURRENT_DATE,INTERVAL -1 DAY))-16*3600
group by  event_type,
task_id 
order by data limit 1

select from_unixtime(UNIX_TIMESTAMP(DATE_SUB(CURRENT_DATE,INTERVAL -1 DAY))-16*3600) as da,UNIX_TIMESTAMP(DATE_SUB(CURRENT_DATE,INTERVAL 5 DAY))-7*3600,UNIX_TIMESTAMP(DATE_SUB(CURRENT_DATE,INTERVAL -1 DAY))-14*3600
select count(id) as data, event_type , task_id from task_log where task_id in(1166,1185,1167,1172,1178,1173,1199,1179,1172,1175,1173,1176,1199,1200,1174,1177,1178,1179,1180,1181,1182,1183) and status = 1 and event_type=2 and created >= UNIX_TIMESTAMP(DATE_SUB(CURRENT_DATE,INTERVAL -1 DAY))-16*3600 group by event_type, task_id 
order by data limit 1;

select count(id) as data, event_type , task_id from task_log where task_id in(1166,1185,1167,1172,1178,1173,1199,1179,1172,1175,1173,1176,1199,1200,1174,1177,1178,1179,1180,1181,1182,1183) and status = 1 and event_type=2 and created >= UNIX_TIMESTAMP(DATE_SUB(CURRENT_DATE,INTERVAL -1 DAY))-16*3600 group by event_type, task_id order by data limit 1;

select  from_unixtime(UNIX_TIMESTAMP(DATE_SUB(CURRENT_DATE,INTERVAL -1 DAY))-16*3600)

select count(id) as num from task_log where task_id = 1185 and event_name = 'incentive_push_action' and created >= UNIX_TIMESTAMP(DATE_SUB(CURRENT_DATE,INTERVAL -1 DAY))-16*3600 



select 
count(id) as num,
event_type,
task_id
from task_log 
where task_id in (1166,1185,1167,1172,1178,1173,1199,1179) 
and status = 1  
and event_type=2
and created >=1689721995 
group by  event_type,
task_id 

(1166, 1167, 1168, 1169, 1170, 1171) 
(1166, 1167, 1168, 1169, 1170, 1171, 1172, 1173, 1174, 1175, 1176, 1177, 1199, 1200, 1178, 1179, 1180, 1181, 1182, 1183)
select 
count(id) as num,
event_type,
task_id
from task_log 
where task_id in(1172, 1173, 1174, 1175, 1176, 1177, 1199, 1200, 1178, 1179, 1180, 1181, 1182, 1183) 
and status = 1  
and event_type=2
and created >=1689523200 
group by  event_type,
task_id 

select * from task_log where task_id = 1260 and  created between UNIX_TIMESTAMP()-86400 and UNIX_TIMESTAMP() 

select count(id) as num from  market_event_log where event_name = 'google_registration' and created between UNIX_TIMESTAMP()-1800 and UNIX_TIMESTAMP()

select count(id) as num,event_name from market_event_log where status = 1 and event_value in ('1005','1006','1007','1008') and created between UNIX_TIMESTAMP()-86400 and UNIX_TIMESTAMP()-101*60 group by event_name

select count(id) as num,event_name from market_event_log where status = 1 and event_value in ('1005','1006','1007','1008') and created between UNIX_TIMESTAMP()-86400 and UNIX_TIMESTAMP() group by event_name

(1172,1178,1173,1199,1179) 
1172,1178

select * from (
select uid,task_id from task_log where task_id in (1172,1178) 
and status = 1  
and event_type=2
and created >=1689721995 order by task_id desc)
as a where a.uid   in (2091257, 1482358, 4727543, 5946982, 9355997, 10500242, 11076946, 11041087, 11233671, 11291988, 11293674, 11867671, 11446648, 11366664, 11383765, 11483942, 10865642, 11509365, 11883693, 11667257, 11014779, 11236859, 11373440, 11228279, 11473881, 11376355, 11584040, 11515762, 11676140, 11630241, 11953540, 4006844, 5066071, 7302278, 6133268, 11172457, 11325748, 11392342, 11313151, 11246940, 11698942, 11530874, 11478040, 11399968, 11674443, 11597773, 11610996, 401986, 1214248, 1575090, 3315766, 1767142, 3717865, 3025948, 2270050, 4353051, 4071445, 5013071, 5693885, 5789850, 3342943, 3816740, 5037450, 4443681, 6057075, 6206745, 4878540, 4627253, 5747986, 7244758, 8816966, 10502372)

select * from task_log where created >=1689721995 and uid = 11373440

select uid,
         JSON_EXTRACT(event_result,
         "$.amount") AS amount,
         JSON_EXTRACT(event_result,
         "$.raise_amount") AS raise_amount,
         JSON_EXTRACT(event_result,
         "$.new_amount") AS new_amount from task_log where uid in(11858880, 11755263, 11713891, 11704369, 11686765, 11575655, 11568684, 11486970, 11456966, 11413079, 11385083, 11384163, 11327981, 11275170, 11268054, 11265854, 11228268, 11226091, 11183876, 11141774, 11039671, 10793361, 10752453, 10742089, 10163591, 10048758, 8245177, 6173051, 5279786, 3432398, 581074) and task_id in (1173,1199,1179) and event_type = 2 and status = 1 AND created >=1688343688
         
       select UNIX_TIMESTAMP(DATE_SUB(CURRENT_DATE,INTERVAL 2 DAY)),UNIX_TIMESTAMP(DATE_SUB(CURRENT_DATE,INTERVAL 1 DAY))
       
       select UNIX_TIMESTAMP(DATE_SUB(CURRENT_DATE,INTERVAL 2 DAY))-7*3600,UNIX_TIMESTAMP(DATE_SUB(CURRENT_DATE,INTERVAL 1 DAY))-7*3600

select * from (
select uid from task_log where task_id in (1172,1178) 
and status = 1  
and event_type=2
and created >=1688343688 order by uid desc)
as a where a.uid not in (3390359, 7491653, 7518370, 11240773, 1383895, 3365754, 3966785, 4040162, 4186060, 7366683, 9021372, 3741791, 3660589, 4232476, 6266550, 5613779, 5823397, 11102876, 11185278, 11123893, 11259198, 11315057, 11048695, 11202088, 11493159, 11479761, 11504793, 11508372, 11523080, 11276367, 11326677, 11364987, 11602595, 11215074, 11621596, 11259582, 575159, 1143172, 2248397, 1924697, 1660190, 2708767, 2324179, 2405895, 2482988, 4337997, 3610176, 3239467, 3675386, 4335684, 2881778, 3655598, 4676363, 5028251, 2527567, 3537956, 3479095, 3809859, 3543976, 3021256, 4145455, 3238156, 3822497, 3657352, 3558892, 6267459, 4453374, 3309267, 7233991, 5427462, 5418892, 7198071, 4259187, 5491052, 4349062, 8271198, 5033974, 4828399, 4348395, 4631761, 6483471, 6136071, 4611371, 8684764, 6702377, 6176859, 7040469, 5557856, 5940682, 7300553, 5742183, 10200683, 7374568)

select * from task_log where uid in(7029653, 5729283, 5726375, 5237381, 4820774, 4733970, 3252595, 2289071) and task_id in (1172,1178) and event_type = 2 and status = 1


SELECT uid,
         JSON_EXTRACT(event_result,
         "$.amount") AS amount,
         JSON_EXTRACT(event_result,
         "$.raise_amount") AS raise_amount,
         JSON_EXTRACT(event_result,
         "$.new_amount") AS new_amount
FROM task_log
WHERE task_id IN (1167)
        AND status = 1
        AND event_type=2
        AND created >=1688343688
ORDER BY  uid desc

select * from task_log where task_id = 1207

select count(task_id) as data from (
select count(id) as num, event_type , task_id from task_log where task_id in(1166,1185,1167,1172,1178,1173,1199,1179,1172,1175,1173,1176,1199,1200,1174,1177,1178,1179,1180,1181,1182,1183) and status = 1 and event_type=2 and created >= UNIX_TIMESTAMP(DATE_SUB(CURRENT_DATE,INTERVAL -1 DAY))-16*3600 group by event_type, task_id) as a

SELECT count(t.task_id) AS data FROM
    (SELECT count(id) AS num,
         event_type ,
         task_id
    FROM task_log
    WHERE task_id in(1166, 1167, 1168, 1169, 1170, 1171, 1172, 1173, 1174, 1175, 1176, 1177, 1199, 1200, 1178, 1179, 1180, 1181, 1182, 1183)
            AND status = 1
            AND event_type=2
            AND created >= UNIX_TIMESTAMP(DATE_SUB(CURRENT_DATE,INTERVAL -1 DAY))-16*3600
    GROUP BY  event_type, task_id) AS t

select from_unixtime(created) as date,uid,event_name,params,event_result from task_log where uid = 1866499 and task_id = 1179 order by id desc;
select * from task_log where uid = 1866499 and task_id = 1179 order by id desc;

select * from task_log where task_id = 1185 and batch_id = 60139 and status = 1 and event_type = 2 AND event_name = 'uangme_coupon_action'
select * from task_log where task_id = 1166 and batch_id = 62076 and status = 1 and event_type = 2



select * from hokuto_msg_topic where message_id in ('1683295187947765760', '1683295187897434112', '1683295187847102464', '1683295187998097408', '1683295188224589824', '1683295188836958208', '1683295189541601280', '1683295189667430400', '1683295189956837376', '1683295189973614592')


select * from hokuto_uangme_token where open_id in ('12174850', '12174133', '12176926', '12173606', '12176839', '12177856', '12179283', 12178734, 12176184, 12179267)


SELECT round((S1.num/S2.total)*100,0) AS data FROM ( SELECT count('id') as num FROM `third_req_res_record` WHERE `type` = '2' and `code` = 'FAIL' ) S1,( SELECT count('id') as total FROM `third_req_res_record` WHERE `type` = '2') S2;、



select *,unix_timestamp(created)

select count(id) as num from third_req_res_record where 

SELECT round((S1.num/S2.total)*100,
        0) AS data
FROM 
    (SELECT count('id') AS num
    FROM `third_req_res_record`
    WHERE `type` = '2'
            AND `code` = 'FAIL' AND created >= unix_timestamp(NOW()-INTERVAL 1 HOUR)) S1,
    (SELECT count('id') AS total
    FROM `third_req_res_record`
    WHERE `type` = '2' AND created >= unix_timestamp(NOW()-INTERVAL 1 HOUR)) S2;
    
    SELECT count('id') AS num
    FROM `third_req_res_record`
    WHERE `type` = '2'
            AND created >= unix_timestamp(NOW()-INTERVAL 7 HOUR)
            
            select unix_timestamp(NOW()-INTERVAL 7 HOUR) as a,NOW()-INTERVAL 1 HOUR as b
            
            select count(id) as num,from_unixtime(created) as a from google_event where event_name = 'high_value_disburse' and created between 1689609600 and 1689696000
            
            	select FROM_UNIXTIME(created,'%Y-%m-%d') as dt ,count(id) as num from google_event  WHERE `event_name` = 'high_value_disburse' and created >= 1688918400 group by dt 

            
            
            SELECT round((S1.num/S2.total)*100, 0) AS data FROM(SELECT count('id') AS num FROM `third_req_res_record` WHERE `type` = '2' AND `code` = 'FAIL' AND created >= unix_timestamp(NOW()-INTERVAL 1 HOUR)) S1, (SELECT count('id') AS total FROM `third_req_res_record` WHERE `type` = '2' AND created >= unix_timestamp(NOW()-INTERVAL 1 HOUR)) S2







select advertising_id,from_unixtime(created) from app_flyer where advertising_id in ("627edadd-66d5-4b50-aba0-5bee9a347617") 




 created between 1687104030 and 1687190430 and

select advertising_id from app_flyer where created between 1687104030 and 1687190430 

SELECT a.id AS uid,
         c.total_amount AS amount,
         d.product_id AS pid
FROM user_info a
LEFT JOIN user_loan b
    ON a.id=b.uid
LEFT JOIN product_account c
    ON a.id=c.uid
LEFT JOIN user_product d
    ON a.id=d.uid
LEFT JOIN user_rc_info u
    ON a.id = u.uid
WHERE a.last_submit_time
        AND a.id=11944148 limit 1
        
        
        SELECT DISTINCT a.uid,a.ctime,a.mtime,
         c.product_id AS pid,a.rtime as real_repayment_time,
         d.total_amount AS amount FROM
    (SELECT *,
         MAX(real_repayment_time) AS rtime
    FROM user_loan_plan
    where uid = 1866499) a
LEFT JOIN user_loan_plan b
    ON a.uid=b.uid
        AND b.ctime>a.rtime
LEFT JOIN user_product c
    ON a.uid=c.uid
LEFT JOIN product_account d
    ON a.uid=d.uid
WHERE  a.uid = 1866499



SELECT DISTINCT a.uid,
         c.product_id AS pid,a.rtime as real_repayment_time,
         d.total_amount AS amount FROM
    (SELECT uid,
         MAX(real_repayment_time) AS rtime
    FROM user_loan_plan
    WHERE real_repayment_time
        BETWEEN UNIX_TIMESTAMP(DATE_SUB(CURRENT_DATE,INTERVAL 5 DAY))-7*3600
            AND UNIX_TIMESTAMP(DATE_SUB(CURRENT_DATE,INTERVAL 4 DAY))-7*3600 and `status`=2
    GROUP BY  uid) a
LEFT JOIN user_loan_plan b
    ON a.uid=b.uid
        AND b.ctime>a.rtime
LEFT JOIN user_product c
    ON a.uid=c.uid
LEFT JOIN product_account d
    ON a.uid=d.uid
WHERE c.is_pass='YES'
        AND c.product_id IN(108,119,116,120,121)
        AND b.id is NULL
        AND d.account_type=2
        AND a.uid % 100 >= 50
GROUP BY  a.uid

select uid,from_unixtime(real_repayment_time) as date,from_unixtime(ctime) as a ,from_unixtime(ctime) as b from user_loan_plan where uid = 12008261
select from_unixtime(last_submit_time) as a from user_info where id = 12008261


select UNIX_TIMESTAMP(DATE_SUB(CURRENT_DATE,INTERVAL 4 DAY))-7*3600,UNIX_TIMESTAMP(DATE_SUB(CURRENT_DATE,INTERVAL 5 DAY))-7*3600

SELECT a.id AS uid
FROM user_info a
LEFT JOIN user_loan b
    ON a.id=b.uid
LEFT JOIN product_account c
    ON a.id=c.uid
LEFT JOIN user_product d
    ON a.id=d.uid
LEFT JOIN user_rc_info u
    ON a.id = u.uid
WHERE a.last_submit_time
    BETWEEN UNIX_TIMESTAMP(DATE_SUB(CURRENT_DATE,INTERVAL 2 DAY)) - 7*3600
        AND UNIX_TIMESTAMP(DATE_SUB(CURRENT_DATE,INTERVAL 1 DAY)) - 7*3600
        AND a.verify_status=2
        AND b.loan_id is NULL
        AND c.account_type=2
        AND d.is_pass='YES'
        AND u.user_risk_lev_v2 in(1,2)
 AND u.id % 100 >= 50
        
GROUP BY  a.id;

 select uid,total_amount as amount from product_account where uid in (11858880, 11755263, 11713891, 11704369, 11686765, 11575655, 11568684, 11486970, 11456966, 11413079, 11385083, 11384163, 11327981, 11275170, 11268054, 11265854, 11228268, 11226091, 11183876, 11141774, 11039671, 10793361, 10752453, 10742089, 10163591, 10048758, 8245177, 6173051, 5279786, 3432398, 581074) AND account_type=2
 
  select uid,total_amount as amount from product_account where uid in (11858880, 11755263, 11713891, 11704369, 11686765, 11575655, 11568684, 11486970, 11456966, 11413079, 11385083, 11384163, 11327981, 11275170, 11268054, 11265854, 11228268, 11226091, 11183876, 11141774, 11039671, 10793361, 10752453, 10742089, 10163591, 10048758, 8245177, 6173051, 5279786, 3432398, 581074) AND account_type=2 order by uid asc

SELECT a.id AS uid, c.total_amount AS amount, d.product_id AS pid FROM user_info a LEFT JOIN user_loan b ON a.id=b.uid LEFT JOIN product_account c ON a.id=c.uid LEFT JOIN user_product d ON a.id=d.uid LEFT JOIN user_rc_info u ON a.id = u.uid WHERE a.last_submit_time BETWEEN UNIX_TIMESTAMP(DATE_SUB(CURRENT_DATE,INTERVAL 3 DAY)) - 7*3600 AND UNIX_TIMESTAMP(DATE_SUB(CURRENT_DATE,INTERVAL 2 DAY)) - 7*3600 AND a.verify_status=2 AND b.loan_id is NULL AND c.account_type=2 AND d.is_pass='YES' AND u.user_risk_lev_v2 in(1,2) AND u.id % 100 >= 50 GROUP BY a.id;



select * from (
SELECT a.id AS uid, c.total_amount AS amount, d.product_id AS pid FROM user_info a LEFT JOIN user_loan b ON a.id=b.uid LEFT JOIN product_account c ON a.id=c.uid LEFT JOIN user_product d ON a.id=d.uid LEFT JOIN user_rc_info u ON a.id = u.uid WHERE a.last_submit_time BETWEEN UNIX_TIMESTAMP(DATE_SUB(CURRENT_DATE,INTERVAL 2 DAY))-7*3600 AND UNIX_TIMESTAMP(DATE_SUB(CURRENT_DATE,INTERVAL 1 DAY))-7*3600 AND a.verify_status=2 AND b.loan_id is NULL AND c.account_type=2 AND d.is_pass='YES' AND u.user_risk_lev_v2 in(1,2) AND u.id % 100 >= 50 GROUP BY a.id) as b where b.uid = 12008261

select * from (
SELECT a.id AS uid, c.total_amount AS amount, d.product_id AS pid FROM user_info a LEFT JOIN user_loan b ON a.id=b.uid LEFT JOIN product_account c ON a.id=c.uid LEFT JOIN user_product d ON a.id=d.uid LEFT JOIN user_rc_info u ON a.id = u.uid WHERE a.last_submit_time BETWEEN UNIX_TIMESTAMP(DATE_SUB(CURRENT_DATE,INTERVAL 2 DAY)) - 7*3600 AND UNIX_TIMESTAMP(DATE_SUB(CURRENT_DATE,INTERVAL 1 DAY)) - 7*3600 AND a.verify_status=2 AND b.loan_id is NULL AND c.account_type=2 AND d.is_pass='YES' AND u.user_risk_lev_v2 in(1,2) AND u.id % 100 >= 50 GROUP BY a.id ) as s where s.uid = 12008261


select * from user_coupon where uid in (4364879, 11711716, 11984030, 11955762, 11972096, 9156132, 11977557, 11987825, 11986962, 11988778, 11986869, 11988697, 11990083, 11988271, 11988263, 11989681, 11989431, 11989845, 11990181, 11990878, 11991817, 11991738, 11991283, 11988777, 11993037, 11990981, 11991001, 11992785, 11990149, 11994620, 11992041, 11991741, 11989847, 11993780, 11992998, 11990720, 11993491, 11994583, 11992935, 11994827, 11991670, 11995102, 11993479) and c_id = 9325

SELECT a.id AS uid,u.user_risk_lev_v2, from_unixtime(u.created) as created,from_unixtime(u.updated) as updated,u.ext,
         c.total_amount AS amount,
         d.product_id AS pid
FROM user_info a
LEFT JOIN user_loan b
    ON a.id=b.uid
LEFT JOIN product_account c
    ON a.id=c.uid
LEFT JOIN user_product d
    ON a.id=d.uid
LEFT JOIN user_rc_info u
    ON a.id = u.uid
WHERE a.last_submit_time
    BETWEEN UNIX_TIMESTAMP(DATE_SUB(CURRENT_DATE,INTERVAL 2 DAY))-7*3600
        AND UNIX_TIMESTAMP(DATE_SUB(CURRENT_DATE,INTERVAL 1 DAY))-7*3600
        AND a.verify_status=2
        AND b.loan_id is NULL
        AND c.account_type=2
        AND d.is_pass='YES'
       
GROUP BY  a.id 

SELECT DISTINCT a.uid,
         c.product_id AS pid,a.rtime as real_repayment_time,
         d.total_amount AS amount FROM
    (SELECT uid,
         MAX(real_repayment_time) AS rtime
    FROM user_loan_plan
    WHERE real_repayment_time
        BETWEEN UNIX_TIMESTAMP(DATE_SUB(CURRENT_DATE,INTERVAL 5 DAY))-7*3600
            AND UNIX_TIMESTAMP(DATE_SUB(CURRENT_DATE,INTERVAL 4 DAY))-7*3600 and `status`=2
    GROUP BY  uid) a
LEFT JOIN user_loan_plan b
    ON a.uid=b.uid
        AND b.ctime>a.rtime
LEFT JOIN user_product c
    ON a.uid=c.uid
LEFT JOIN product_account d
    ON a.uid=d.uid
WHERE c.is_pass='YES'
        AND c.product_id IN(108,119,116,120,121)
        AND b.id is NULL
        AND d.account_type=2
        AND a.uid % 100 < 50
GROUP BY  a.uid

select from_unixtime(UNIX_TIMESTAMP(DATE_SUB(CURRENT_DATE,INTERVAL 2 DAY))-7*3600),from_unixtime(UNIX_TIMESTAMP(DATE_SUB(CURRENT_DATE,INTERVAL 1 DAY))-7*3600)

SELECT DISTINCT a.uid,
         c.product_id AS pid,a.rtime as real_repayment_time, 
         d.total_amount AS amount FROM
    (SELECT uid,
         MAX(real_repayment_time) AS rtime
    FROM user_loan_plan
    WHERE real_repayment_time
        BETWEEN UNIX_TIMESTAMP(DATE_SUB(CURRENT_DATE,INTERVAL 3 DAY))-7*3600
            AND UNIX_TIMESTAMP(DATE_SUB(CURRENT_DATE,INTERVAL 2 DAY))-7*3600 and `status`=2
    GROUP BY  uid) a
LEFT JOIN user_loan_plan b
    ON a.uid=b.uid
        AND b.ctime>a.rtime
LEFT JOIN user_product c
    ON a.uid=c.uid
LEFT JOIN product_account d
    ON a.uid=d.uid
WHERE c.is_pass='YES'
        AND c.product_id IN(108,119,116,120,121)
        AND b.id is NULL
        AND d.account_type=2
        AND a.uid % 100 < 50
GROUP BY  a.uid

SELECT a.id AS uid,
         c.total_amount AS amount,
         d.product_id AS pid
FROM user_info a
LEFT JOIN user_loan b
    ON a.id=b.uid
LEFT JOIN product_account c
    ON a.id=c.uid
LEFT JOIN user_product d
    ON a.id=d.uid
LEFT JOIN user_rc_info u
    ON a.id = u.uid
WHERE a.last_submit_time
    BETWEEN UNIX_TIMESTAMP(DATE_SUB(CURRENT_DATE,INTERVAL 4 DAY))-7*3600
        AND UNIX_TIMESTAMP(DATE_SUB(CURRENT_DATE,INTERVAL 3 DAY))-7*3600
        AND a.verify_status=2
        AND b.loan_id is NULL
        AND c.account_type=2
        AND d.is_pass='YES'

        AND u.id % 100 >= 50
GROUP BY  a.id

select distinct a.uid,a.user_level from user_fdc_score as a left join user_info as b on a.uid = b.id where a.user_level in (2,3) and a.created >= UNIX_TIMESTAMP() - 3*3600 and a.created <= UNIX_TIMESTAMP() and b.verify_status<=0;

select distinct a.uid from user_fdc_score as a left join user_info as b on a.uid = b.id where a.user_level =2 and a.created >= UNIX_TIMESTAMP() - 3*3600 and a.created <= UNIX_TIMESTAMP() and b.verify_status<=0;
select ul.uid,ul.id,ul.loan_id,ul.status,ulp.status as s,ul.loan_product_id,ul.money,from_unixtime(ul.ctime) as cdate,from_unixtime(ul.mtime) as mdate from user_loan ul left join user_loan_plan as ulp on ul.uid = ulp.uid where ul.uid = 658710
order by cdate desc

select * from user_rc_info where uid in (10156955,10160970,10165161,10461570,1065730,10807951,10843371,10850355,1183556,1242362,1358600,1392170,1418195,151291,1584661,1737154,1773160,1984456,2074184,2092996,2230551,2381266,2548280,2567360,2691796,2743770,2794756,2823665,3041573,3122376,3236053,3331770,3439693,3751230,385070,4003182,4405865,4413830,4491293,4759291,4980976,5184852,5429081,5493771,6187766,6304081,6363966,6898840,6927171,6930075,7022394,7097306,7431665,7524592,7568261,76280,7858684,8006454,8156866,8503302,8582672,8826340,903256,9134356,9172684,9414362,9457950) order by updated desc


select distinct a.uid from user_fdc_score as a left join user_info as b on a.uid = b.id where a.user_level = 2 and a.created >= UNIX_TIMESTAMP() - 1800 and a.created <= UNIX_TIMESTAMP() and b.verify_status<=0;
select distinct a.uid from user_fdc_score as a left join user_info as b on a.uid = b.id where a.user_level = 3 and a.created >= UNIX_TIMESTAMP() - 1800 and a.created <= UNIX_TIMESTAMP() and b.verify_status<=0;
select distinct a.uid from user_fdc_score as a left join user_info as b on a.uid = b.id where a.user_level = 2 and a.created >= UNIX_TIMESTAMP() - 3*3600 and a.created <= UNIX_TIMESTAMP() and b.verify_status<=0;
select distinct a.uid from user_fdc_score as a left join user_info as b on a.uid = b.id where a.user_level = 3 and a.created >= UNIX_TIMESTAMP() - 3*3600 and a.created <= UNIX_TIMESTAMP() and b.verify_status<=0;

select distinct a.uid from user_fdc_score as a left join user_info as b on a.uid = b.id where a.user_level = 3 and a.created between UNIX_TIMESTAMP() - 2400 and  UNIX_TIMESTAMP()-1800 and b.verify_status<=0;

select distinct a.uid from user_fdc_score as a left join user_info as b on a.uid = b.id where a.user_level = 3 and a.created between UNIX_TIMESTAMP() - 3*3600-2400 and  UNIX_TIMESTAMP()-1800 and b.verify_status<=0;

SELECT DISTINCT ufs.uid FROM user_fdc_score AS ufs LEFT JOIN user_info AS ui ON ufs.uid = ui.id WHERE ufs.created BETWEEN UNIX_TIMESTAMP() - 600 AND UNIX_TIMESTAMP() - 300 AND ui.verify_status = 0 AND ui.first_input_base_info_time = 0;
user_fdc_score

SELECT DISTINCT ufs.uid FROM user_info AS ui LEFT JOIN user_fdc_score AS ufs ON ui.id = ufs.uid WHERE ufs.created BETWEEN UNIX_TIMESTAMP() - 600 AND UNIX_TIMESTAMP() - 300 AND ui.verify_status = 0 AND ui.first_input_base_info_time = 0 AND ufs.uid % 100 <= 25;

SELECT  id AS uid FROM user_info WHERE created BETWEEN UNIX_TIMESTAMP() - 2400 AND UNIX_TIMESTAMP() - 300 AND verify_status = 0 AND first_input_base_info_time = 0 

AND id % 100 <= 25;

SELECT id AS uid FROM user_info WHERE created BETWEEN UNIX_TIMESTAMP() - 600 AND UNIX_TIMESTAMP() - 300 AND verify_status = 0 AND first_input_base_info_time = 0 AND id % 100 <= 25; 
select UNIX_TIMESTAMP()
SELECT id AS uid
FROM user_info
WHERE created
    BETWEEN UNIX_TIMESTAMP() - 1800
        AND UNIX_TIMESTAMP() - 900
        AND verify_status = 0
        AND first_input_base_info_time = 0
        AND id % 100 <= 25; 

select  UNIX_TIMESTAMP() - 1800

between UNIX_TIMESTAMP() - 1800 and UNIX_TIMESTAMP()
between UNIX_TIMESTAMP() - 3*3600 and UNIX_TIMESTAMP()

SELECT id AS uid
FROM user_info
WHERE created
    BETWEEN UNIX_TIMESTAMP() - 1200
        AND UNIX_TIMESTAMP() - 600
        AND verify_status = 0
        AND first_input_base_info_time = 0
        AND id % 100 <= 24; 



select * from (
SELECT DISTINCT a.uid,
         c.product_id AS pid,a.rtime as real_repayment_time,
         d.total_amount AS amount FROM
    (SELECT uid,
         MAX(real_repayment_time) AS rtime
    FROM user_loan_plan
    WHERE real_repayment_time
        BETWEEN UNIX_TIMESTAMP(DATE_SUB(CURRENT_DATE,INTERVAL 5 DAY))-7*3600
            AND UNIX_TIMESTAMP(DATE_SUB(CURRENT_DATE,INTERVAL 4 DAY))-7*3600 and `status`=2
    GROUP BY  uid) a
LEFT JOIN user_loan_plan b
    ON a.uid=b.uid
        AND b.ctime>a.rtime
LEFT JOIN user_product c
    ON a.uid=c.uid
LEFT JOIN product_account d
    ON a.uid=d.uid
WHERE c.is_pass='YES'
        AND c.product_id IN(108,119,116,120,121)
        AND b.id is NULL
        AND d.account_type=2
        AND a.uid % 100 >= 40
GROUP BY  a.uid) as a where uid = 11373440


SELECT DISTINCT a.uid,pa.total_amount as amount
FROM user_fdc_score AS a
LEFT JOIN user_info AS b
    ON a.uid = b.id
    left join product_account as pa
    ON a.uid = pa.uid
WHERE a.user_level = 3
        AND a.created
    BETWEEN UNIX_TIMESTAMP() - 3*3600 - 2400
        AND UNIX_TIMESTAMP() - 1800
        AND b.verify_status<=0; 

        INSERT INTO `market_strategy` (`strategy_name`, `project`, `filters`, `action`, `status`, `is_test`)
VALUES
	('注册成功', 'uangme', NULL, '[{\"label\":\"market_event执行器\",\"name\":\"market_event_action\"}]', 1, 2);


INSERT INTO `market_event_rule` (`strategy_id`, `action`, `filters`, `template`, `event_name`, `event_value`, `source`, `status`, `ext`)
VALUES
	(4, 'market_event_action', '', '{\"value\": \"{{$value}}\"}', 'google_loan_success', '1004', 'google', 1, NULL);


INSERT INTO `coupon_temporary_quota`(`name`, `desc`, `amount`, `total_count`, `user_limit`, `start_time`, `end_time`, `is_combine`, `is_open`, `loan_product_id`, `created`, `updated`) VALUES ('Kupon Kenaikan Kuota', 'tambahan <font color=\'#246CFD\'><b>Rp 80.000</b></font> untuk meningkatkan batas kredit Anda!', 80000, 999999, 1, 1689868800, 1918742400, 2, 1, '',1689907503,1689907503); 
INSERT INTO `coupon_temporary_quota`(`name`, `desc`, `amount`, `total_count`, `user_limit`, `start_time`, `end_time`, `is_combine`, `is_open`, `loan_product_id`, `created`, `updated`) VALUES ('Kupon Kenaikan Kuota', 'tambahan <font color=\'#246CFD\'><b>Rp 40.000</b></font> untuk meningkatkan batas kredit Anda!', 40000, 999999, 1, 1689868800, 1918742400, 2, 1, '',1689907503,1689907503); 



INSERT INTO `coupon_temporary_quota`(`name`, `desc`, `amount`, `total_count`, `user_limit`, `start_time`, `end_time`, `is_combine`, `is_open`, `loan_product_id`, `created`, `updated`) VALUES('Kupon Kenaikan Kuota', 'tambahan <font color=\'#246CFD\'><b>Rp 80.000</b></font> untuk meningkatkan batas kredit Anda!', 80000, 999999, 1, 1689868800, 1918742400, 2, 1, '',1689907503,1689907503), ('Kupon Kenaikan Kuota', 'tambahan <font color=\'#246CFD\'><b>Rp 40.000</b></font> untuk meningkatkan batas kredit Anda!', 40000, 999999, 1, 1689868800, 1918742400, 2, 1, '',1689907503,1689907503); 

INSERT INTO `market_strategy` (`strategy_name`, `project`, `filters`, `action`, `status`, `is_test`, `created`, `updated`)
VALUES
	('贷前V3评级打点事件1-4', 'uangme', NULL, '[{\"label\":\"market_event执行器\",\"name\":\"market_event_action\"}]', 1, 2, 0, 0);
	
INSERT INTO `market_event_rule` ( `strategy_id`, `action`, `filters`, `template`, `event_name`, `event_value`, `source`, `status`, `ext`, `created`, `updated`)
VALUES
	( 22, 'market_event_action', '', '', 'google_v3_1_3', '1005', 'google', 1, NULL, 0, 0);
	
	INSERT INTO `market_event_rule` ( `strategy_id`, `action`, `filters`, `template`, `event_name`, `event_value`, `source`, `status`, `ext`, `created`, `updated`)
VALUES
	( 22, 'market_event_action', '', '', 'google_v3_1_4', '1006', 'google', 1, NULL, 0, 0);
	
	INSERT INTO `market_event_rule` ( `strategy_id`, `action`, `filters`, `template`, `event_name`, `event_value`, `source`, `status`, `ext`, `created`, `updated`)
VALUES
	( 22, 'market_event_action', '', '', 'appflyer_v3_1_3', '1007', 'appflyer', 1, NULL, 0, 0);
	
	INSERT INTO `market_event_rule` ( `strategy_id`, `action`, `filters`, `template`, `event_name`, `event_value`, `source`, `status`, `ext`, `created`, `updated`)
VALUES
	( 22, 'market_event_action', '', '', 'appflyer_v3_1_4', '1008', 'appflyer', 1, NULL, 0, 0);
	
	
	INSERT INTO `market_strategy` ( `strategy_name`, `project`, `filters`, `action`, `status`, `is_test`, `created`, `updated`)
VALUES
	('贷前V3评级打点事件', 'uangme', NULL, '[{\"label\":\"market_event执行器\",\"name\":\"market_event_action\"}]', 1, 2, 0, 0);


SELECT DISTINCT a.uid
FROM user_fdc_score AS a
LEFT JOIN user_info AS b
    ON a.uid = b.id
    left join device as d
    On d.
WHERE a.id = null
        AND a.created
    BETWEEN UNIX_TIMESTAMP() - 3*3600 - 2400
        AND UNIX_TIMESTAMP() - 1800
        AND b.verify_status<=0; 
        
        
SELECT DISTINCT ui.id as uid
FROM user_info AS ui
LEFT JOIN user_fdc_score AS ufs
    ON ui.id = ufs.uid
    left join device as d
    On ui.id = d.u_id
WHERE ufs.id = null and d.user_level = 1
        AND ui.created
    BETWEEN UNIX_TIMESTAMP() - 3*3600 - 2400
        AND UNIX_TIMESTAMP() - 1800
        AND ui.verify_status<=0; 
        
        
        
        SELECT DISTINCT ui.id AS uid
FROM user_info AS ui
LEFT JOIN user_fdc_score AS ufs
    ON ui.id = ufs.uid
LEFT JOIN device AS d
    ON ui.id = d.u_id
WHERE ufs.id is null
        AND d.user_level = 1
        AND ui.created
    BETWEEN UNIX_TIMESTAMP() - 3*3600 - 2400
        AND UNIX_TIMESTAMP() - 1800
        AND ui.verify_status<=0; 
        
           SELECT DISTINCT ui.id AS uid
FROM user_info AS ui,user_fdc_score AS ufs
LEFT JOIN device AS d
    ON ui.id = d.u_id
WHERE ufs.id = null
        AND d.user_level = 1
        AND ui.created
    BETWEEN UNIX_TIMESTAMP() - 3*3600 - 2400
        AND UNIX_TIMESTAMP() - 1800
        AND ui.verify_status<=0;
        
        
        INSERT INTO `coupon_temporary_quota`(`name`, `desc`, `amount`, `total_count`, `user_limit`, `start_time`, `end_time`, `is_combine`, `is_open`, `loan_product_id`) VALUES ('Kupon Kenaikan Kuota', 'tambahan <font color=\'#246CFD\'><b>Rp 80.000</b></font> untuk meningkatkan batas kredit Anda!', 80000, 999999, 1, 1689868800, 1729440000, 2, 1, ''); 





恢复数据：

UPDATE `task_log` SET `created` = '1687742068' WHERE `task_id` in (1166, 1167, 1168) and created=0 and status = 1 and event_type= 2;
UPDATE `task_log` SET `created` = '1687742068' WHERE `task_id` in (1166, 1167, 1168) and created=0;
UPDATE `task_log` SET `created` = '1687742068' WHERE `task_id` in (1169, 1170, 1171) and created=0;


UPDATE `task_log` SET `created` = '1687742068' WHERE `task_id` in (1172,1173,1174,1199) and created=0 and status = 1 and event_type = 2;
UPDATE `task_log` SET `created` = '1687742068' WHERE `task_id` in (1175,1176,1177,1200) and created=0 and status = 1 and event_type = 2;
UPDATE `task_log` SET `created` = '1687742068' WHERE `task_id` in (1178,1179,1180) and created=0 and status = 1 and event_type = 2;
UPDATE `task_log` SET `created` = '1687742068' WHERE `task_id` in (1181,1182,1183) and created=0 and status = 1 and event_type = 2;
UPDATE `task_log` SET `created` = '1687742068' WHERE `task_id` in (1172,1173,1174,1199,1175,1176,1177,1200,1178,1179,1180,1181,1182,1183) and created=0 and status = 1 and event_type = 2;


(1166,
1167,
1168,
1169,
1170,
1171,
1172,
1173,
1174,
1175,
1176,
1177,
1199,
1200,
1178,
1179,
1180,
1181,
1182,
1183)




(7029653,
5729283,
5726375,
5237381,
4820774,
4733970,
3252595,
2289071)


(3390359,
7491653,
7518370,
11240773,
1383895,
3365754,
3966785,
4040162,
4186060,
7366683,
9021372,
3741791,
3660589,
4232476,
6266550,
5613779,
5823397,
11102876,
11185278,
11123893,
11259198,
11315057,
11048695,
11202088,
11493159,
11479761,
11504793,
11508372,
11523080,
11276367,
11326677,
11364987,
11602595,
11215074,
11621596,
11259582,
575159,
1143172,
2248397,
1924697,
1660190,
2708767,
2324179,
2405895,
2482988,
4337997,
3610176,
3239467,
3675386,
4335684,
2881778,
3655598,
4676363,
5028251,
2527567,
3537956,
3479095,
3809859,
3543976,
3021256,
4145455,
3238156,
3822497,
3657352,
3558892,
6267459,
4453374,
3309267,
7233991,
5427462,
5418892,
7198071,
4259187,
5491052,
4349062,
8271198,
5033974,
4828399,
4348395,
4631761,
6483471,
6136071,
4611371,
8684764,
6702377,
6176859,
7040469,
5557856,
5940682,
7300553,
5742183,
10200683,
7374568)



-7*3600

SELECT a.id AS uid,
         c.total_amount AS amount,
         d.product_id AS pid
FROM user_info a
LEFT JOIN user_loan b
    ON a.id=b.uid
LEFT JOIN product_account c
    ON a.id=c.uid
LEFT JOIN user_product d
    ON a.id=d.uid
LEFT JOIN user_rc_info u
    ON a.id = u.uid
WHERE a.last_submit_time
    BETWEEN UNIX_TIMESTAMP(DATE_SUB(CURRENT_DATE,INTERVAL 2 DAY))-7
        AND UNIX_TIMESTAMP(DATE_SUB(CURRENT_DATE,INTERVAL 1 DAY))
        AND a.verify_status=2
        AND b.loan_id is NULL
        AND c.account_type=2
        AND d.is_pass='YES'
        AND u.user_risk_lev_v2 in(1,2)
        AND u.id % 100 >= 50
GROUP BY  a.id




 {
      "label": "激励推送执行器",
      "name": "incentive_push_action",
      "params": {
          "amount": {
              "label": "用户当前额度",
              "field_type": "tag_input",
              "type": "",
              "value": ""
          },
          "strategy_push": {
              "label": "是否推送激励",
              "field_type": "select",
              "options": [
                  {"label": "否","value":0},
                  {"label": "是","value":1}
              ],
              "type": "int",
              "value": 0
          },
          "strategy_repeat": {
              "label": "激励类型",
              "field_type": "select",
              "options": [
                  {"label": "首贷","value":1},
                  {"label": "复贷","value":2}
              ],
              "type": "int",
              "value": 1
          },
          "strategy_day": {
              "label": "推送激励策略",
              "field_type": "select",
              "options": [
                  {"label": "无","value":0},
                  {"label": "D1","value":1},
                  {"label": "D2","value":2},
                  {"label": "D3","value":3},
                  {"label": "D4","value":4},
                  {"label": "D5","value":5}
              ],
              "type": "int",
              "value": 0
          }
      }
 },

type BindEmpQuotaArgs struct {
	Uid         int64  `form:"uid" binding:"required"`
	CouponId    int64  `form:"coupon_id"`
	TotalAmount int64  `form:"total_amount"`
	StartTime   int64  `form:"start_time" binding:"required"`
	EndTime     int64  `form:"end_time" binding:"required"`
	Status      int8   `form:"status"`
	IsRead      int8   `form:"is_read"`
	Source      string `form:"source"`
}

(4364879,
11711716,
11984030,
11955762,
11972096,
9156132,
11977557,
11987825,
11986962,
11988778,
11986869,
11988697,
11990083,
11988271,
11988263,
11989681,
11989431,
11989845,
11990181,
11990878,
11991817,
11991738,
11991283,
11988777,
11993037,
11990981,
11991001,
11992785,
11990149,
11994620,
11992041,
11991741,
11989847,
11993780,
11992998,
11990720,
11993491,
11994583,
11992935,
11994827,
11991670,
11995102,
11993479)

 internal_api_ip: 
        - "121.78.231.21/32"
        - "127.0.0.1/32"
        - "10.50.100.40/32"
        - "10.84.100.68/32"
        - "103.69.155.90/32"
        - "103.69.155.91/32"
        - "103.69.155.92/32"
        - "103.69.155.93/32"
        - "182.18.33.196/32"
        - "182.18.33.197/32"
        - "182.18.33.198/32"
        - "182.18.33.199/32"
        - "202.66.35.243/32"
        - "202.66.35.244/32"
        - "202.66.35.245/32"
        - "202.66.35.246/32"
        - "13.228.49.19/32"
        - "103.69.155.80/32"
        - "137.59.101.109/32"
        - "208.90.121.182/32"
        - "103.220.76.131/32"
        - "103.220.76.132/32"
        - "121.78.231.21/32"
        - "121.78.231.22/32"
        - "147.139.185.188/32"
        - "10.185.9.133/32"
        - "103.150.185.227/32"
        - "182.253.84.191/32"
        - "182.253.127.230/32"




CREATE TABLE `data_statistics` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '记录id',
  `task_id` varchar(200)  NOT NULL DEFAULT '0' COMMENT '任务id',
  `push_success_num` int(11) unsigned NOT NULL DEFAULT '0' COMMENT 'push的成功数',
  `push_topic_success_num` int(11) unsigned NOT NULL DEFAULT '0' COMMENT 'hokuto_msg_topic的push的成功数',
  `push_report_click_num` int(11) unsigned NOT NULL DEFAULT '0' COMMENT 'push的点击数',
  `push_report_reach_num` int(11) unsigned NOT NULL DEFAULT '0' COMMENT 'push的到达数',
  `push_order_loan_num` int(11) unsigned NOT NULL DEFAULT '0' COMMENT 'push的用户的下单数',
  `status` tinyint(3) unsigned NOT NULL DEFAULT 1 COMMENT '数据状态: 2:无效,1:有效',
  `created` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '任务创建时间',
  `updated` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '任务修改时间',
  PRIMARY KEY (`id`),
  KEY `IDX_created` (`created`),
  KEY `IDX_task_id` (`task_id`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='数据统计表';




我的验证逻辑

1.1 luban.task_log记录表
 a)、获取被过滤的用户的uid
 b)、获取成功提额的用户uid
 1.2 crm_feature.rc_outdata_repayment_result校验数据
 a)、校验被过滤的用户的uid
 b)、校验被成功提额的用户的uid的数据



smsContent = strings.Replace(smsContent, "{{amount}}", fmt.Sprint(Amount), -1)
smsContent = strings.Replace(smsContent, "{{new_amount}}", fmt.Sprint(result["new_amount"]), -1)
pushContent = strings.Replace(pushContent, "{{amount}}", fmt.Sprint(Amount), -1)
pushContent = strings.Replace(pushContent, "{{new_amount}}", fmt.Sprint(result["new_amount"]), -1)

2023年06月25日 星期日
select count(id) as num ,event_type,task_id from task_log where task_id in (1178,1179,1180,1181,1182,1183) and status = 1  and created >= 1687648480 group by  event_type,  task_id 
select count(id) as num ,event_type,task_id from task_log where task_id in (1172,1173,1174,1175,1176,1177,1199,1200) and status = 1  and created >= 1687648480 group by  event_type,  task_id 
select count(id) as num ,event_type,task_id from task_log where task_id in (1166,1185,1167,1172,1178,1173,1199,1179) and status = 1  and created >= 1687648480 group by  event_type,  task_id 


(11890019,
11911001,
11235585,
11911707,
11904781,
11912083,
11913391,
11897445,
11905002,
11910503,
11897618,
11911304,
11912905,
11909400,
11911505,
9889486,
11315932)


(11858880,
11755263,
11713891,
11704369,
11686765,
11575655,
11568684,
11486970,
11456966,
11413079,
11385083,
11384163,
11327981,
11275170,
11268054,
11265854,
11228268,
11226091,
11183876,
11141774,
11039671,
10793361,
10752453,
10742089,
10163591,
10048758,
8245177,
6173051,
5279786,
3432398,
581074)


SELECT t1.id,
         t1.uid,
         t1.ctime,
         t1.mtime ,
         substr(t1.ctime,
         1,
         10) AS c_date ,
         substr(t1.mtime,
         1,
         10) AS m_date ,
         JSON_EXTRACT(t1.data,
         "$.exter_first_trans_to_apply_days") AS exter_first_trans_to_apply_days ,
         JSON_EXTRACT(t1.data,
         "$.dz_level_short_v2") AS dz_level_short_v2
FROM crm_feature.rc_outdata_repayment_result AS t1
WHERE t1.uid IN(392777, 557693, 453657, 478583, 735299, 568351, 498486, 767183, 664355, 542291, 523762, 448753, 650480, 709578, 675660, 844780, 752175, 724967, 869497, 1136777, 630060) 



SELECT DISTINCT a.uid,
         c.product_id AS pid,
         d.total_amount AS amount
FROM 
    (SELECT uid,
         MAX(real_repayment_time) AS rtime
    FROM user_loan_plan
    WHERE real_repayment_time
        BETWEEN UNIX_TIMESTAMP(DATE_SUB(CURRENT_DATE,INTERVAL 4 DAY))
            AND UNIX_TIMESTAMP(DATE_SUB(CURRENT_DATE,INTERVAL 3 DAY))and `status`=2
    GROUP BY  uid) a
LEFT JOIN user_loan_plan b
    ON a.uid=b.uid
        AND b.ctime>a.rtime
LEFT JOIN user_product c
    ON a.uid=c.uid
LEFT JOIN product_account d
    ON a.uid=d.uid
WHERE c.is_pass='YES'
        AND c.product_id IN(108,119,116,120,121)
        AND b.id is NULL
        AND d.account_type=2
        AND a.uid % 100 >= 50
GROUP BY  a.uid

func (obj *uangmeRepo) DealData(ctx context.Context, taskId string, startTime int64, endTime int64) (bool, error) {
	num, err_task_log := obj.lubanDb.Master().Table("task_log").Where("created BETWEEN ? AND ?", startTime, endTime).In("task_id", strArr).FindAndCount(ctx, dataTaskLog1)
	win, err = obj.lubanDb.Master().Table("task_log").Select("batch_id").Where("created BETWEEN ? AND ?", startTime, endTime).In("task_id", strArr).Get(ctx, dataTaskLog2)
	err = obj.lubanDb.Master().Table("task_log").Select("batch_id").Where("created BETWEEN ? AND ?", startTime, endTime).In("task_id", strArr).GroupBy("batch_id").Find(ctx, data2)
	num2, errs1 := obj.cmpushInDb.Master().Table("hokuto_msg_topic").Where("state = ? AND success =  ?", 1, 1).In("push_id", strArr2).FindAndCount(ctx, hokutoMsgTopicData2)
	num3, errs2 := obj.cmpushInDb.Master().Table("hokuto_msg_report").Where("status >= ", 0).In("push_id", strArr2).FindAndCount(ctx, hokutoMsgTopicData2)
	num4, errs3 := obj.cmpushInDb.Master().Table("hokuto_msg_report").Where("status >= ", 1).In("push_id", strArr2).FindAndCount(ctx, hokutoMsgTopicData2)
	affect1, err1 := obj.db.Master().Table("data_statistics").Where("created = ? AND task_id = ?", dataStatics.Created, dataStatics.TaskId).Update(ctx, dataStaticsUpdate)
}










for i, c := range ups {
		switch c.ProductId {
		case 119:
			if i > 0 {
				remove := new(model.UserProduct)
				remove = ups[0]
				newUpss := make([]*model.UserProduct, 0)
				newUpss = append(ups[:0], ups[(0+1):]...)
				newUpss[0] = ups[i]
				newUpss = append(newUpss, remove)
				ups = newUpss
			}
		default:
			break
		}
	}

  {"sms_channel_type":"0","sms_content":"yugjgku"}





  SELECT DISTINCT a.uid,
         c.product_id AS pid,
         d.total_amount AS amount FROM
    (SELECT uid,
         MAX(real_repayment_time) AS rtime
    FROM user_loan_plan
    WHERE real_repayment_time
        BETWEEN UNIX_TIMESTAMP(DATE_SUB(CURRENT_DATE,INTERVAL 4 DAY))
            AND UNIX_TIMESTAMP(DATE_SUB(CURRENT_DATE,INTERVAL 3 DAY))and `status`=2
    GROUP BY  uid) a
LEFT JOIN user_loan_plan b
    ON a.uid=b.uid
        AND b.ctime>a.rtime
LEFT JOIN user_product c
    ON a.uid=c.uid
LEFT JOIN product_account d
    ON a.uid=d.uid
WHERE c.is_pass='YES'
        AND c.product_id IN(108,119,116,120,121)
        AND b.id is NULL
        AND d.account_type=2
        AND a.uid % 100 >= 50
GROUP BY  a.uid


select uid,
         JSON_EXTRACT(event_result,
         "$.amount") AS amount,
         JSON_EXTRACT(event_result,
         "$.raise_amount") AS raise_amount,
         JSON_EXTRACT(event_result,
         "$.new_amount") AS new_amount from task_log where uid in(11858880, 11755263, 11713891, 11704369, 11686765, 11575655, 11568684, 11486970, 11456966, 11413079, 11385083, 11384163, 11327981, 11275170, 11268054, 11265854, 11228268, 11226091, 11183876, 11141774, 11039671, 10793361, 10752453, 10742089, 10163591, 10048758, 8245177, 6173051, 5279786, 3432398, 581074) and task_id in (1173,1199,1179) and event_type = 2 and status = 1 AND created >=1688343688




