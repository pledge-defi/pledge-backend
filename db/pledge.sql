-- MySQL dump 10.13  Distrib 5.7.34, for Linux (x86_64)
--
-- Host: localhost    Database: pledge_v22
-- ------------------------------------------------------
-- Server version	5.7.34-log

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `admin`
--

DROP TABLE IF EXISTS `admin`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `admin`
(
    `user_id`  int(11) NOT NULL AUTO_INCREMENT,
    `name`     varchar(100) NOT NULL,
    `password` varchar(100) NOT NULL,
    PRIMARY KEY (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `admin`
--

LOCK
TABLES `admin` WRITE;
/*!40000 ALTER TABLE `admin` DISABLE KEYS */;
INSERT INTO `admin`
VALUES (1, 'admin', '$2b$12$V4n31JgG1ybX1NuAUgp6keYw5eAyCrNQ5TAl/FyRtgBtF9/uf1Q6S');
/*!40000 ALTER TABLE `admin` ENABLE KEYS */;
UNLOCK
TABLES;

--
-- Table structure for table `multi_sign`
--

DROP TABLE IF EXISTS `multi_sign`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `multi_sign`
(
    `id`                 int(10) unsigned NOT NULL AUTO_INCREMENT,
    `sp_name`            varchar(255) DEFAULT NULL,
    `sp_token`           varchar(255) DEFAULT NULL,
    `jp_name`            varchar(255) DEFAULT NULL,
    `jp_token`           varchar(255) DEFAULT NULL,
    `sp_address`         varchar(255) DEFAULT NULL,
    `jp_address`         varchar(255) DEFAULT NULL,
    `sp_hash`            varchar(255) DEFAULT NULL,
    `jp_hash`            varchar(255) DEFAULT NULL,
    `multi_sign_account` varchar(255) DEFAULT NULL,
    `chain_id`           int(10) DEFAULT NULL,
    `created_at`         date         DEFAULT NULL,
    `updated_at`         date         DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `multi_sign`
--

LOCK
TABLES `multi_sign` WRITE;
/*!40000 ALTER TABLE `multi_sign` DISABLE KEYS */;
/*!40000 ALTER TABLE `multi_sign` ENABLE KEYS */;
UNLOCK
TABLES;

--
-- Table structure for table `poolbases`
--

DROP TABLE IF EXISTS `poolbases`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `poolbases`
(
    `id`                       int(11) NOT NULL AUTO_INCREMENT,
    `settle_time`              varchar(100) DEFAULT NULL,
    `end_time`                 varchar(100) DEFAULT NULL,
    `interest_rate`            varchar(100) DEFAULT NULL,
    `max_supply`               varchar(100) DEFAULT NULL,
    `lend_supply`              varchar(100) DEFAULT NULL,
    `borrow_supply`            varchar(100) DEFAULT NULL,
    `martgage_rate`            varchar(100) DEFAULT NULL,
    `lend_token`               varchar(100) DEFAULT NULL,
    `borrow_token`             varchar(100) DEFAULT NULL,
    `state`                    varchar(100) DEFAULT NULL,
    `jp_coin`                  varchar(100) DEFAULT NULL,
    `sp_coin`                  varchar(100) DEFAULT NULL,
    `auto_liquidate_threshold` varchar(100) DEFAULT NULL,
    `created_at`               datetime     DEFAULT NULL,
    `updated_at`               datetime     DEFAULT NULL,
    `pool_id`                  int(11) DEFAULT NULL,
    `borrow_token_info`        json         DEFAULT NULL,
    `lend_token_info`          json         DEFAULT NULL,
    `chain_id`                 varchar(20)  DEFAULT '56',
    `lend_token_symbol`        varchar(100) DEFAULT NULL,
    `borrow_token_symbol`      varchar(100) DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1213 DEFAULT CHARSET=utf8 COMMENT='poolbase';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `poolbases`
--

LOCK
TABLES `poolbases` WRITE;
/*!40000 ALTER TABLE `poolbases` DISABLE KEYS */;
INSERT INTO `poolbases`
VALUES (1198, '1644824159', '1644824759', '10000000', '100000000000000000000000', '200000000000000000000', '0',
        '200000000', '0xE676Dcd74f44023b95E0E2C6436C97991A7497DA', '0xB5514a4FA9dDBb48C3DE215Bc9e52d9fCe2D8658', '4',
        '0x02Fc7E9E3A2E07f31ecEc6bD8b89144D4051cFC6', '0xD3588D695519a5088177C6cd215d61b398bbC9B0', '1000000',
        '2022-02-18 16:08:27', '2022-02-21 12:33:12', 1, '{
    \"borrowFee\": \"0\", \"tokenLogo\": \"https://v2-backend.pledger.finance/storage/img/BTC.png\", \"tokenName\": \"BTC\", \"tokenPrice\": \"3911641941261\"}',
        '{
          \"lendFee\": \"0\", \"tokenLogo\": \"https://v2-backend.pledger.finance/storage/img/BUSD.png\", \"tokenName\": \"BUSD\", \"tokenPrice\": \"100029930\"}',
        '97', 'BUSD', 'BTC'),
       (1199, '1644825059', '1644825359', '10000000', '100000000000000000000000', '200000000000000000000',
        '100000000000000', '200000000', '0xE676Dcd74f44023b95E0E2C6436C97991A7497DA',
        '0xB5514a4FA9dDBb48C3DE215Bc9e52d9fCe2D8658', '2', '0x8FE38F33Da9038f8959b7f7c0E613276937e0276',
        '0xCdFE3fE5C0A41BC003223636d9c22f9bA796cE2F', '1000000', '2022-02-18 16:08:28', '2022-02-21 12:33:13', 2, '{
         \"borrowFee\": \"0\", \"tokenLogo\": \"https://v2-backend.pledger.finance/storage/img/BTC.png\", \"tokenName\": \"BTC\", \"tokenPrice\": \"3911641941261\"}',
        '{
          \"lendFee\": \"0\", \"tokenLogo\": \"https://v2-backend.pledger.finance/storage/img/BUSD.png\", \"tokenName\": \"BUSD\", \"tokenPrice\": \"100029930\"}',
        '97', 'BUSD', 'BTC'),
       (1200, '1645322400', '1645495200', '20000000', '100000000000000000000000', '2300000000000000000000',
        '44199067531432704', '200000000', '0xE676Dcd74f44023b95E0E2C6436C97991A7497DA',
        '0xB5514a4FA9dDBb48C3DE215Bc9e52d9fCe2D8658', '1', '0xD5E39E3dA2399f49b14A66f6763F8C4fcAACf097',
        '0x51D6955BB212e79f6bF3bDCDc31910012Fa164B0', '20000000', '2022-02-18 16:08:29', '2022-02-21 12:33:13', 3, '{
         \"borrowFee\": \"0\", \"tokenLogo\": \"https://v2-backend.pledger.finance/storage/img/BTC.png\", \"tokenName\": \"BTC\", \"tokenPrice\": \"3911641941261\"}',
        '{
          \"lendFee\": \"0\", \"tokenLogo\": \"https://v2-backend.pledger.finance/storage/img/BUSD.png\", \"tokenName\": \"BUSD\", \"tokenPrice\": \"100029930\"}',
        '97', 'BUSD', 'BTC'),
       (1201, '1643558400', '1643817600', '5000000', '100000000000000000000000', '20000000000000000000',
        '506810558736854', '200000000', '0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56',
        '0x7130d2A12B9BCbFAe4f2634d864A1Ee1Ce3Ead9c', '2', '0xEC446F606bF543511080031A1d839842056eDF24',
        '0x3f9b13Bc64D6Da2b56A20d83755Cf2D06430ff7D', '20000000', '2022-02-18 16:08:30', '2022-02-21 12:34:44', 1, '{
         \"borrowFee\": \"250000\", \"tokenLogo\": \"https://v2-backend.pledger.finance/storage/img/BTC.png\", \"tokenName\": \"BTC\", \"tokenPrice\": \"3899244000000\"}',
        '{
          \"lendFee\": \"250000\", \"tokenLogo\": \"https://v2-backend.pledger.finance/storage/img/BUSD.png\", \"tokenName\": \"BUSD\", \"tokenPrice\": \"100002724\"}',
        '56', 'BUSD', 'BTC'),
       (1202, '1644005970', '1644265178', '4000000', '100000000000000000000000', '65000000000000000000',
        '1070946941727608', '200000000', '0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56',
        '0x7130d2A12B9BCbFAe4f2634d864A1Ee1Ce3Ead9c', '3', '0x8Dbe02eB53c409756Ca250137996F1ceaFAf3557',
        '0xaFAbf67BD0c1FB875F17DF9063328982A2538a7c', '100000000', '2022-02-18 16:08:31', '2022-02-21 12:34:44', 2, '{
         \"borrowFee\": \"250000\", \"tokenLogo\": \"https://v2-backend.pledger.finance/storage/img/BTC.png\", \"tokenName\": \"BTC\", \"tokenPrice\": \"3899244000000\"}',
        '{
          \"lendFee\": \"250000\", \"tokenLogo\": \"https://v2-backend.pledger.finance/storage/img/BUSD.png\", \"tokenName\": \"BUSD\", \"tokenPrice\": \"100002724\"}',
        '56', 'BUSD', 'BTC'),
       (1203, '1643402775', '1643402968', '5000000', '100000000000000000000000', '10000000000000000000',
        '962887635827337', '200000000', '0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56',
        '0x7130d2A12B9BCbFAe4f2634d864A1Ee1Ce3Ead9c', '2', '0x43C6195cFBEfDFfb98e8aA043c7AEeE83EdB1e24',
        '0x08622e15fa321b66D55846EDAa1f0912432ea6ad', '30000000', '2022-02-18 16:08:32', '2022-02-21 12:34:45', 3, '{
         \"borrowFee\": \"250000\", \"tokenLogo\": \"https://v2-backend.pledger.finance/storage/img/BTC.png\", \"tokenName\": \"BTC\", \"tokenPrice\": \"3899244000000\"}',
        '{
          \"lendFee\": \"250000\", \"tokenLogo\": \"https://v2-backend.pledger.finance/storage/img/BUSD.png\", \"tokenName\": \"BUSD\", \"tokenPrice\": \"100002724\"}',
        '56', 'BUSD', 'BTC'),
       (1204, '1644018452', '1644094063', '6000000', '10000000000000000000000', '51000000000000000000',
        '1395716938127436', '200000000', '0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56',
        '0x7130d2A12B9BCbFAe4f2634d864A1Ee1Ce3Ead9c', '2', '0x26B99638c525296B04a0150044A9e23D61b95a3b',
        '0x2FaC08e5Ff93CEdb5AA2678c35B885A23454B585', '20000000', '2022-02-18 16:08:32', '2022-02-21 12:34:46', 4, '{
         \"borrowFee\": \"250000\", \"tokenLogo\": \"https://v2-backend.pledger.finance/storage/img/BTC.png\", \"tokenName\": \"BTC\", \"tokenPrice\": \"3899244000000\"}',
        '{
          \"lendFee\": \"250000\", \"tokenLogo\": \"https://v2-backend.pledger.finance/storage/img/BUSD.png\", \"tokenName\": \"BUSD\", \"tokenPrice\": \"100002724\"}',
        '56', 'BUSD', 'BTC'),
       (1205, '1644014387', '1644360003', '7000000', '100000000000000000000000', '1000000000000000000', '0',
        '200000000', '0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56', '0x7130d2A12B9BCbFAe4f2634d864A1Ee1Ce3Ead9c', '4',
        '0x4dd305aCC09bc180Ac14fA24F0895BE6f7065A44', '0x4D63B40FaC1254AeF2Ed43e05D02c6DaF215208A', '20000000',
        '2022-02-18 16:08:33', '2022-02-21 12:34:46', 5, '{
         \"borrowFee\": \"250000\", \"tokenLogo\": \"https://v2-backend.pledger.finance/storage/img/BTC.png\", \"tokenName\": \"BTC\", \"tokenPrice\": \"3899244000000\"}',
        '{
          \"lendFee\": \"250000\", \"tokenLogo\": \"https://v2-backend.pledger.finance/storage/img/BUSD.png\", \"tokenName\": \"BUSD\", \"tokenPrice\": \"100002724\"}',
        '56', 'BUSD', 'BTC'),
       (1206, '1644638402', '1645243208', '4000000', '10000000000000000000000', '30000000000000000000',
        '960262149375210', '200000000', '0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56',
        '0x7130d2A12B9BCbFAe4f2634d864A1Ee1Ce3Ead9c', '2', '0xe268c022E80F42DdFDfcdBdf5BB0Df0b77e53CcC',
        '0xB1A31db22655c74dc0886C2E6E44Bf2cC3539Ce1', '20000000', '2022-02-18 16:08:34', '2022-02-21 12:34:47', 6, '{
         \"borrowFee\": \"250000\", \"tokenLogo\": \"https://v2-backend.pledger.finance/storage/img/BTC.png\", \"tokenName\": \"BTC\", \"tokenPrice\": \"3899244000000\"}',
        '{
          \"lendFee\": \"250000\", \"tokenLogo\": \"https://v2-backend.pledger.finance/storage/img/BUSD.png\", \"tokenName\": \"BUSD\", \"tokenPrice\": \"100002724\"}',
        '56', 'BUSD', 'BTC'),
       (1207, '1645244011', '1645848818', '3000000', '10000000000000000000000', '40000000000000000000', '0',
        '200000000', '0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56', '0x7130d2A12B9BCbFAe4f2634d864A1Ee1Ce3Ead9c', '4',
        '0xDa2e5b8484e8A760615027d8721458D794f64bbB', '0xaa3621852264F91Be9d9450D8f5B06C97DDDEF47', '20000000',
        '2022-02-18 16:08:35', '2022-02-21 12:34:47', 7, '{
         \"borrowFee\": \"250000\", \"tokenLogo\": \"https://v2-backend.pledger.finance/storage/img/BTC.png\", \"tokenName\": \"BTC\", \"tokenPrice\": \"3899244000000\"}',
        '{
          \"lendFee\": \"250000\", \"tokenLogo\": \"https://v2-backend.pledger.finance/storage/img/BUSD.png\", \"tokenName\": \"BUSD\", \"tokenPrice\": \"100002724\"}',
        '56', 'BUSD', 'BTC'),
       (1208, '1645244304', '1645935511', '5000000', '10000000000000000000000', '10000000000000000000', '0',
        '200000000', '0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56', '0x7130d2A12B9BCbFAe4f2634d864A1Ee1Ce3Ead9c', '4',
        '0x80C4b556Df7CEEd97bE4F396dfaB68244e7877AC', '0xaD3fC9593b66617Cbaf9154deF567657302eE907', '20000000',
        '2022-02-18 16:08:36', '2022-02-21 12:34:48', 8, '{
         \"borrowFee\": \"250000\", \"tokenLogo\": \"https://v2-backend.pledger.finance/storage/img/BTC.png\", \"tokenName\": \"BTC\", \"tokenPrice\": \"3899244000000\"}',
        '{
          \"lendFee\": \"250000\", \"tokenLogo\": \"https://v2-backend.pledger.finance/storage/img/BUSD.png\", \"tokenName\": \"BUSD\", \"tokenPrice\": \"100002724\"}',
        '56', 'BUSD', 'BTC'),
       (1209, '1645516800', '1645689600', '10000000', '300000000000000000000000', '100000000000000000000', '0',
        '200000000', '0xE676Dcd74f44023b95E0E2C6436C97991A7497DA', '0xB5514a4FA9dDBb48C3DE215Bc9e52d9fCe2D8658', '0',
        '0x5B88dc80d1336fb99E69872fd25bC36E7295Ec45', '0x8ADD9b86ef46Db612348C7fC8fC789070f36a9c6', '20000000',
        '2022-02-18 16:43:25', '2022-02-21 12:33:14', 4, '{
         \"borrowFee\": \"0\", \"tokenLogo\": \"https://v2-backend.pledger.finance/storage/img/BTC.png\", \"tokenName\": \"BTC\", \"tokenPrice\": \"3911641941261\"}',
        '{
          \"lendFee\": \"0\", \"tokenLogo\": \"https://v2-backend.pledger.finance/storage/img/BUSD.png\", \"tokenName\": \"BUSD\", \"tokenPrice\": \"100029930\"}',
        '97', 'BUSD', 'BTC'),
       (1210, '1645892800', '1646497607', '5000000', '10000000000000000000000', '20000000000000000000',
        '499304444511538', '200000000', '0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56',
        '0x7130d2A12B9BCbFAe4f2634d864A1Ee1Ce3Ead9c', '0', '0xE4EBB886e94f6299AbCf84E1981cDe1AEEB44612',
        '0xCFC7d3F53f7ea055c24f130C5bb9933cd4de7921', '20000000', '2022-02-21 11:07:30', '2022-02-21 12:34:49', 9, '{
         \"borrowFee\": \"250000\", \"tokenLogo\": \"https://v2-backend.pledger.finance/storage/img/BTC.png\", \"tokenName\": \"BTC\", \"tokenPrice\": \"3899244000000\"}',
        '{
          \"lendFee\": \"250000\", \"tokenLogo\": \"https://v2-backend.pledger.finance/storage/img/BUSD.png\", \"tokenName\": \"BUSD\", \"tokenPrice\": \"100002724\"}',
        '56', 'BUSD', 'BTC'),
       (1211, '1645893533', '1646498337', '3000000', '10000000000000000000000', '0', '0', '200000000',
        '0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56', '0x7130d2A12B9BCbFAe4f2634d864A1Ee1Ce3Ead9c', '0',
        '0x51A8Ff282c57b4a131dcCb4C01e790a0D90b064A', '0xc92Dc9ddf7263d69562CC1934973e479f1F37895', '100000000',
        '2022-02-21 11:07:31', '2022-02-21 12:34:49', 10, '{
         \"borrowFee\": \"250000\", \"tokenLogo\": \"https://v2-backend.pledger.finance/storage/img/BTC.png\", \"tokenName\": \"BTC\", \"tokenPrice\": \"3899244000000\"}',
        '{
          \"lendFee\": \"250000\", \"tokenLogo\": \"https://v2-backend.pledger.finance/storage/img/BUSD.png\", \"tokenName\": \"BUSD\", \"tokenPrice\": \"100002724\"}',
        '56', 'BUSD', 'BTC'),
       (1212, '1645893787', '1647704604', '4000000', '10000000000000000000000', '0', '0', '200000000',
        '0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56', '0x7130d2A12B9BCbFAe4f2634d864A1Ee1Ce3Ead9c', '0',
        '0x0048b307f070A1AE0D9148d707f82C058D86b53A', '0x98CC7541CcB15eCDDA0f71820b4EE11B5C0B2a83', '100000000',
        '2022-02-21 11:07:32', '2022-02-21 12:34:50', 11, '{
         \"borrowFee\": \"250000\", \"tokenLogo\": \"https://v2-backend.pledger.finance/storage/img/BTC.png\", \"tokenName\": \"BTC\", \"tokenPrice\": \"3899244000000\"}',
        '{
          \"lendFee\": \"250000\", \"tokenLogo\": \"https://v2-backend.pledger.finance/storage/img/BUSD.png\", \"tokenName\": \"BUSD\", \"tokenPrice\": \"100002724\"}',
        '56', 'BUSD', 'BTC');
/*!40000 ALTER TABLE `poolbases` ENABLE KEYS */;
UNLOCK
TABLES;

--
-- Table structure for table `pooldata`
--

DROP TABLE IF EXISTS `pooldata`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `pooldata`
(
    `settle_amount_lend`       varchar(100) DEFAULT NULL,
    `settle_amount_borrow`     varchar(100) DEFAULT NULL,
    `finish_amount_lend`       varchar(100) DEFAULT NULL,
    `finish_amount_borrow`     varchar(100) DEFAULT NULL,
    `liquidation_amoun_lend`   varchar(100) DEFAULT NULL,
    `liquidation_amoun_borrow` varchar(100) DEFAULT NULL,
    `updated_at`               datetime     DEFAULT NULL,
    `created_at`               datetime     DEFAULT NULL,
    `id`                       int(11) NOT NULL AUTO_INCREMENT,
    `chain_id`                 varchar(20)  DEFAULT '56',
    `pool_id`                  varchar(50)  DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1219 DEFAULT CHARSET=utf8 COMMENT='pooldata';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `pooldata`
--

LOCK
TABLES `pooldata` WRITE;
/*!40000 ALTER TABLE `pooldata` DISABLE KEYS */;
INSERT INTO `pooldata`
VALUES ('200000000000000000000', '0', '0', '0', '0', '0', '2022-02-21 12:34:40', '2022-02-18 16:08:27', 1204, '97',
        '1'),
       ('2116489638425690052', '100000000000000', '2116491651207336194', '47667589312225', '0', '0',
        '2022-02-21 12:34:40', '2022-02-18 16:08:28', 1205, '97', '2'),
       ('883929609006842180605', '44199067531432704', '0', '0', '0', '0', '2022-02-21 12:34:41', '2022-02-18 16:08:29',
        1206, '97', '3'),
       ('9622480365307393626', '506810558736854', '9701569167251517196', '246056897508601', '0', '0',
        '2022-02-21 12:34:44', '2022-02-18 16:08:31', 1207, '56', '1'),
       ('21784681613977652294', '1070946941727608', '0', '0', '21963739279576028628', '527106498724037',
        '2022-02-21 12:34:45', '2022-02-18 16:08:31', 1208, '56', '2'),
       ('10000000000000000000', '531338857291495', '10000061100000000000', '263993253351156', '0', '0',
        '2022-02-21 12:34:45', '2022-02-18 16:08:32', 1209, '56', '3'),
       ('28574755457698650835', '1395716938127436', '28643266291384029120', '700751333610913', '0', '0',
        '2022-02-21 12:34:46', '2022-02-18 16:08:33', 1210, '56', '4'),
       ('1000000000000000000', '0', '0', '0', '0', '0', '2022-02-21 12:34:46', '2022-02-18 16:08:34', 1211, '56', '5'),
       ('20332950849176712243', '960262149375210', '20722901670458952508', '439767062265160', '0', '0',
        '2022-02-21 12:34:47', '2022-02-18 16:08:34', 1212, '56', '6'),
       ('40000000000000000000', '0', '0', '0', '0', '0', '2022-02-21 12:34:48', '2022-02-18 16:08:35', 1213, '56', '7'),
       ('10000000000000000000', '0', '0', '0', '0', '0', '2022-02-21 12:34:48', '2022-02-18 16:08:36', 1214, '56', '8'),
       ('0', '0', '0', '0', '0', '0', '2022-02-21 12:34:43', '2022-02-18 16:43:26', 1215, '97', '4'),
       ('0', '0', '0', '0', '0', '0', '2022-02-21 12:34:49', '2022-02-21 11:07:30', 1216, '56', '9'),
       ('0', '0', '0', '0', '0', '0', '2022-02-21 12:34:50', '2022-02-21 11:07:31', 1217, '56', '10'),
       ('0', '0', '0', '0', '0', '0', '2022-02-21 12:34:50', '2022-02-21 11:07:32', 1218, '56', '11');
/*!40000 ALTER TABLE `pooldata` ENABLE KEYS */;
UNLOCK
TABLES;

--
-- Table structure for table `token_info`
--

DROP TABLE IF EXISTS `token_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `token_info`
(
    `id`             int(10) unsigned NOT NULL AUTO_INCREMENT,
    `symbol`         varchar(100) DEFAULT NULL,
    `logo`           varchar(150) DEFAULT NULL,
    `price`          varchar(50)  DEFAULT NULL,
    `token`          varchar(100) DEFAULT NULL,
    `chain_id`       varchar(20)  DEFAULT '56',
    `abi_file_exist` int(2) unsigned DEFAULT '0',
    `created_at`     datetime     DEFAULT NULL,
    `updated_at`     datetime     DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=246 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `token_info`
--

LOCK
TABLES `token_info` WRITE;
/*!40000 ALTER TABLE `token_info` DISABLE KEYS */;
INSERT INTO `token_info`
VALUES (233, 'BTC', 'https://v2-backend.pledger.finance/storage/img/BTC.png', '3899319533584',
        '0xB5514a4FA9dDBb48C3DE215Bc9e52d9fCe2D8658', '97', 0, '2022-02-18 16:08:27', '2022-02-21 12:34:59'),
       (234, 'BUSD', 'https://v2-backend.pledger.finance/storage/img/BUSD.png', '100029930',
        '0xE676Dcd74f44023b95E0E2C6436C97991A7497DA', '97', 0, '2022-02-18 16:08:27', '2022-02-21 12:34:59'),
       (235, 'BTC', 'https://v2-backend.pledger.finance/storage/img/BTC.png', '3900207791936',
        '0x7130d2A12B9BCbFAe4f2634d864A1Ee1Ce3Ead9c', '56', 0, '2022-02-18 16:08:30', '2022-02-21 12:34:59'),
       (236, 'BUSD', 'https://v2-backend.pledger.finance/storage/img/BUSD.png', '100002724',
        '0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56', '56', 1, '2022-02-18 16:08:30', '2022-02-21 12:34:59'),
       (237, 'PLGR', 'https://v2-backend.pledger.finance/storage/img/PLGR.png', '0',
        '0x6Aa91CbfE045f9D154050226fCc830ddbA886CED', '56', 0, '2022-02-18 16:08:45', '2022-02-21 12:34:59'),
       (238, 'ETH', 'https://v2-backend.pledger.finance/storage/img/ETH.png', '0',
        '0x2170ed0880ac9a755fd29b2688956bd959f933f8', '56', 0, '2022-02-18 16:08:45', '2022-02-21 12:34:59'),
       (239, 'CAKE', 'https://v2-backend.pledger.finance/storage/img/CAKE.png', '762359964',
        '0x0e09fabb73bd3ade0a17ecc321fd13a19e81ce82', '56', 0, '2022-02-18 16:08:46', '2022-02-21 12:34:59'),
       (240, 'CAKE', 'https://v2-backend.pledger.finance/storage/img/CAKE.png', '500000000',
        '0xEAEd08168a2D34Ae2B9ea1c1f920E0BC00F9fA67', '97', 0, '2022-02-18 16:08:46', '2022-02-21 12:34:59'),
       (241, 'BNB', 'https://v2-backend.pledger.finance/storage/img/BNB.png', '38844888841',
        '0x0000000000000000000000000000000000000000', '97', 0, '2022-02-18 16:08:46', '2022-02-21 12:34:59'),
       (242, 'BNB', 'https://v2-backend.pledger.finance/storage/img/BNB.png', '38913919138',
        '0x0000000000000000000000000000000000000000', '56', 0, '2022-02-18 16:08:46', '2022-02-21 12:34:59'),
       (243, 'DAI', 'https://v2-backend.pledger.finance/storage/img/DAI.png', '99557039',
        '0x490BC3FCc845d37C1686044Cd2d6589585DE9B8B', '97', 0, '2022-02-18 16:08:46', '2022-02-21 12:34:59'),
       (244, 'DAI', 'https://v2-backend.pledger.finance/storage/img/DAI.png', '99999114',
        '0x1AF3F329e8BE154074D8769D1FFa4eE058B1DBc3', '56', 1, '2022-02-18 16:08:46', '2022-02-21 12:34:59'),
       (245, 'USDT', 'https://v2-backend.pledger.finance/storage/img/USDT.png', '100044886',
        '0x55d398326f99059ff775485246999027b3197955', '56', 1, '2022-02-18 16:08:46', '2022-02-21 12:34:59');
/*!40000 ALTER TABLE `token_info` ENABLE KEYS */;
UNLOCK
TABLES;

--
-- Dumping events for database 'pledge_v22'
--

--
-- Dumping routines for database 'pledge_v22'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-02-21 19:05:59
