CREATE DATABASE  IF NOT EXISTS `MyDB` /*!40100 DEFAULT CHARACTER SET latin1 */;
USE `MyDB`;
--
-- Table structure for table `MessageBoard`
--

DROP TABLE IF EXISTS `MessageBoard`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `MessageBoard` (
  `topicName` varchar(100) NOT NULL,
  `content` varchar(10000) DEFAULT NULL,
  `timestamp` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
