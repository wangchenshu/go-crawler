# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.7.28)
# Database: atomy
# Generation Time: 2020-01-05 11:03:03 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table centers
# ------------------------------------------------------------

DROP TABLE IF EXISTS `centers`;

CREATE TABLE `centers` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `address` varchar(255) DEFAULT NULL,
  `phone` varchar(100) DEFAULT NULL,
  `date` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

LOCK TABLES `centers` WRITE;
/*!40000 ALTER TABLE `centers` DISABLE KEYS */;

INSERT INTO `centers` (`id`, `name`, `address`, `phone`, `date`)
VALUES
	(1,'台中黎明','臺中市 南屯區 黎明路一段812-10號2樓','04-23861507','2014-01-08'),
	(2,'台北中路','台北市 中山區 民權西路19號5樓','02-25850211','2014-01-08'),
	(3,'台北站前','臺北市 中正區 忠孝西路一段50號24樓之1','02-23317882','2014-01-08'),
	(4,'高雄中山','高雄市 前鎮區 中山二路91號20樓之1','07-3347815','2014-01-08'),
	(5,'新北信義','新北市 板橋區 信義路103號1樓','02-29540808','2014-01-08'),
	(6,'台南喜樂多','臺南市 東區 東門路三段253號7樓之1','06-2891739','2014-01-08'),
	(7,'高雄齊納福','高雄市苓雅區 中正二路20號5樓','07-9686688','2014-01-08'),
	(8,'中壢千展','桃園市 中壢區 中央東路88號21樓之1','03-4266090','2014-01-08'),
	(9,'台北忠孝聯合','臺北市 中正區 羅斯福路二段70號4樓之1','02-23218809','2014-01-09'),
	(10,'桃園富豪','桃園市 桃園區 中正路1015號5樓','03-3566300','2014-02-05'),
	(11,'台中冠軍','臺中市 北區 忠明路428號','04-22026497','2014-10-01'),
	(12,'台中朝馬','臺中市 西屯區 台灣大道三段556巷1弄32號','04-24519028','2015-04-16'),
	(13,'屏東信望愛','屏東縣屏東市 和平路354號4樓','08-7661909','2015-07-01'),
	(14,'竹北耀威','新竹縣 竹北市 新泰路92號4樓','03-5535019','2015-08-03'),
	(15,'台中中港','台中市 西屯區 台灣大道三段220號','04-23128770','2015-08-03'),
	(16,'台中1491','台中市 北區 崇德路一段579號6樓之一','04-22349191','2015-08-17'),
	(17,'桃園尚倫','桃園市 桃園區 建國路101號2樓','03-3665226','2015-11-02'),
	(18,'新北文化','新北市 板橋區 文化路一段333號3樓之1','02-22581303','2015-11-02'),
	(19,'南投千峰','南投縣 草屯鎮 新生路64號','04-92359292','2016-01-04'),
	(20,'嘉義六合','嘉義市西區 劉厝路219號','05-2833033','2016-03-16'),
	(21,'台中潭子','台中市潭子區 中山路二段178號','04-25355400','2016-03-16'),
	(22,'桃園龍潭','桃園市龍潭區 中豐路313號','03-4990628','2016-04-01'),
	(23,'草屯勝利','南投縣草屯鎮 成功路一段142號','04-92330142','2016-04-01'),
	(24,'新北板橋','新北市板橋區 長江路三段62號','02-82588067','2016-05-16'),
	(25,'台中豐原','台中市 豐原區 圓環北路一段226-2號3樓','04-25270788','2016-09-01'),
	(26,'台東太平洋','台東縣台東市 四維路一段377號','08-9343849','2016-09-16'),
	(27,'新竹幸福','新竹市 中正路20號3樓之2','03-5350251','2016-10-17'),
	(28,'花蓮成功','花蓮縣吉安鄉 自強路278號','03-8535086','2016-11-01'),
	(29,'新北卓越','新北市新店區 民權路88號2樓之4','02-22188088','2016-11-01'),
	(30,'宜蘭禾家康','宜蘭縣 羅東鎮 傳藝路三段216號2樓','03-9531532','2016-11-01'),
	(31,'高雄鳳山','高雄市 鳳山區 八德路199號2樓','07-7677619','2016-12-17'),
	(32,'苗栗向陽','苗栗縣 竹南鎮 三平路1號','03-7631549','2017-05-16'),
	(33,'雲林斗六','雲林縣 斗六市 明德北路二段407號','05-5335066','2017-06-16'),
	(34,'高雄博愛','高雄市 三民區 博愛一路55號10樓','07-3230681','2017-07-03'),
	(35,'台南和緯','台南市 北區 和緯路五段172號','06-2803446','2017-08-01'),
	(36,'台南府城','台南市 中西區 府前路一段283號7樓之5','06-2139196','2017-08-01'),
	(37,'高雄全球','高雄市 鼓山區 美術東五路66號5樓','07-5227788','2017-08-16'),
	(38,'新竹誠愛','新竹市東區 忠孝路182號3樓','03-5733278','2017-09-18'),
	(39,'內壢叮咚','桃園市中壢區 興農路125號4樓','03-4621024','2017-11-16'),
	(40,'岡山翔富','高雄市岡山區 文賢路37號','07-6211388','2017-12-18'),
	(41,'嘉義陽光','嘉義市東區 市宅街10之1號2樓1','05-2277211','2018-04-02'),
	(42,'新竹日光','新竹市東區 關新二街90號6樓','03-5780301','2018-04-02'),
	(43,'基隆大月','基隆市安樂區 安一路166號1樓','02-24287803','2018-04-16'),
	(44,'汐止齊心','新北市汐止區 信義路1號2樓C8','02-26919905','2018-05-02'),
	(45,'台南佳里','台南市佳里區 公園路666號2樓','06-7237217','2018-07-16'),
	(46,'三重百兆','新北市三重區 福隆路9號2樓','02-82870949','2018-08-16'),
	(47,'高雄世光','高雄市左營區 重立路856號1樓','07-3502349','2018-08-16'),
	(48,'台南永康','台南市永康區 中山南路688號1樓','06-2010678','2018-10-01'),
	(49,'屏東天意','屏東縣屏東市 建南路8號3樓','08-7560286','2018-11-16'),
	(50,'台北富綠地','台北市信義區 基隆路一段155號5樓之9','02-27678285','2019-01-02'),
	(51,'高雄皇冠','高雄市三民區 建工路613號4樓','07-3928638','2019-01-02'),
	(52,'南桃園正善','桃園市桃園區 國強一街418號','03-3601059','2019-01-16'),
	(53,'埔里成堡','南投縣埔里鎮 中正路260號','04-92900918','2019-02-01'),
	(54,'高雄崇尚','高雄市大寮區 中正路74號','07-7836660','2019-02-01'),
	(55,'彰化金馬','彰化縣彰化市 永安街379號3樓','04-7357248','2019-03-04'),
	(56,'北大啄木鳥','新北市三峽區 大義路29號','02-26728208','2019-03-18'),
	(57,'高雄右昌','高雄市楠梓區 加昌路624號','07-3645508','2019-03-18'),
	(58,'屏東艾綠願','屏東縣鹽埔鄉 新二村維新路42-19號','08-7930099','2019-04-16'),
	(59,'高雄楠梓','高雄市楠梓區 建楠路139號5F','07-3536666','2019-05-16'),
	(60,'員林愛加倍','員林市大饒里 柳橋路一段653號','04-8332915','2019-05-16'),
	(61,'旗山','高雄市旗山區 復新東街25號','07-6623141','2019-05-16'),
	(62,'台中大里得勝','台中市大里區 新芳路31號','04-24822325','2019-06-03'),
	(63,'台中樂富','台中市南區 國光路365巷19之1號','04-22859255','2019-06-17'),
	(64,'高雄三民','高雄市三民區 大昌二路266號10樓','07-3836303','2019-06-17'),
	(65,'高雄永豐','高雄市前鎮區 永豐路36號3樓','07-7211236','2019-06-17'),
	(66,'歸仁大發','台南市歸仁區 大順街109號','06-2300059','2019-07-01'),
	(67,'大台中','台中市北屯區 北屯路388之8號','04-22471389','2019-07-01'),
	(68,'大甲順天','台中市大甲區 順天里文武路135號3樓','04-26763639','2019-07-16'),
	(69,'台北大安夢想','台北市大安區 信義路三段166巷4號1樓','02-23252244','2019-07-16'),
	(70,'台中沙鹿真愛','台中市沙鹿區 光榮街90巷9號','04-26627767','2019-08-16'),
	(71,'台南新營','台南市新營區 健康路87號','06-6565789','2019-09-02'),
	(72,'高雄橋頭富盛','高雄市橋頭區 成功北路180號','07-6130968','2019-09-16'),
	(73,'台南真善美','台南市永康區 中華路12號六樓之3','06-3120716','2020-01-02');

/*!40000 ALTER TABLE `centers` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
