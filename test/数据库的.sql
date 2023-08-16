
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
and created >=1688343688 
group by  event_type,
task_id 

(1166, 1167, 1168, 1169, 1170, 1171) 
(1166, 1167, 1168, 1169, 1170, 1171, 1172, 1173, 1174, 1175, 1176, 1177, 1199, 1200, 1178, 1179, 1180, 1181, 1182, 1183)
select 
count(id) as num,
event_type,
task_id
from task_log 
where task_id in(1166, 1167, 1168, 1169, 1170, 1171, 1172, 1173, 1174, 1175, 1176, 1177, 1199, 1200, 1178, 1179, 1180, 1181, 1182, 1183)
and status = 1  
and event_type=2
and created >=1689032687 
group by  event_type,
task_id 

select * from task_log where task_id = 1170 and  created between 1688259606 and 1688346006 








select * from (
select uid from task_log where task_id in (1173,1199,1179) 
and status = 1  
and event_type=2
and created >=1688343688 order by uid desc)
as a where a.uid not in (1614751,4544469,4549594,5635853,5833351,6654956,7375487,11243255,11616893,653692,1364784,3801765,5364360,7334274,10870590,11459154,11514675,3566254,3356985,4551261,6904178,6596761,9004361,10397187,10269762,10898958,11091867,10911693,10819784,10930064,11176883,11039258,10956186,11067959,10982490,11077278,11608752,11579268,10962364,11183497,11507564,11221677,11156054,10932252,11473259,11145657,11339377,11313154,11728288,11891187,11520572,11283597,11420272,11302767,11463457,11387151,11595476,11486870,11640065,173961,474958,512793,1001297,1700093,881880,1698195,1599175,487855,3063879,853987,709578,1755274,2570758,3551658,750651,2654289,1611952,2020657,1540067,967258,2077796,3922268,3053379,2516693,2847285,3598252,2543376,4569482,1737172,4497890,3600482,1666669,2672294,6380966,2326577,5002076,4663859,4448789,4825267,3409065,4428260,6469363,3404177,4800669,4951457,4025558,4230158,7317882,3711756,5490869,8066360,4242962,5566491,3606181,5520459,8662655,8109755,3673085,5182459,7528688,6918453,8001387,8279589,7139369,7215070,6313956,6516998,8027583,9105182,8280571,8452184,10599958,6677966,7649072,8822279,7054354,8673667,7343858,9466888,8938389,5215390,10327551,9701665,7711153,8969691,9715161,5554283,10472450,7745268,9386663,9045474,9951374,10103752,8397883,8835964,10355699,9173068,8539256,7374465,8653571,10672273,9521575,9259694,8913060,9361856,9620398,10017262,10622170)

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
