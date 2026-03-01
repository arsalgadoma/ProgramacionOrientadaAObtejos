-- MySQL dump 10.13  Distrib 8.0.42, for Win64 (x86_64)
--
-- Host: localhost    Database: proyectofinal
-- ------------------------------------------------------
-- Server version	8.0.42

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `orden`
--

DROP TABLE IF EXISTS `orden`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `orden` (
  `OrdenID` int NOT NULL AUTO_INCREMENT,
  `Total` float NOT NULL,
  `Fecha` datetime NOT NULL,
  `UsuarioID` int NOT NULL,
  `Detalle` json DEFAULT NULL,
  PRIMARY KEY (`OrdenID`),
  KEY `fk_UsuaurioID_idx` (`UsuarioID`),
  CONSTRAINT `fk_usuario` FOREIGN KEY (`UsuarioID`) REFERENCES `usuario` (`UsuarioID`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `orden`
--

LOCK TABLES `orden` WRITE;
/*!40000 ALTER TABLE `orden` DISABLE KEYS */;
INSERT INTO `orden` VALUES (9,20,'2026-02-28 23:35:19',1,NULL),(10,30,'2026-02-28 23:47:10',1,'[{\"nombre\": \"Zapatos\", \"cantidad\": 1, \"subtotal\": 20, \"productoId\": 1, \"precioUnitario\": 20}, {\"nombre\": \"Camista\", \"cantidad\": 1, \"subtotal\": 10, \"productoId\": 2, \"precioUnitario\": 10}]'),(11,10,'2026-02-28 23:49:24',1,'[{\"nombre\": \"Camista\", \"cantidad\": 1, \"subtotal\": 10, \"productoId\": 2, \"precioUnitario\": 10}]'),(12,10,'2026-02-28 23:51:34',1,'[{\"nombre\": \"Camista\", \"cantidad\": 1, \"subtotal\": 10, \"productoId\": 2, \"precioUnitario\": 10}]'),(13,10,'2026-02-28 23:53:28',1,'[{\"nombre\": \"Camista\", \"cantidad\": 1, \"subtotal\": 10, \"productoId\": 2, \"precioUnitario\": 10}]'),(14,160,'2026-03-01 00:25:47',1,'[{\"nombre\": \"Camista\", \"cantidad\": 1, \"subtotal\": 10, \"productoId\": 2, \"precioUnitario\": 10}, {\"nombre\": \"Telefono\", \"cantidad\": 1, \"subtotal\": 150, \"productoId\": 3, \"precioUnitario\": 150}]'),(15,150,'2026-03-01 00:26:42',1,'[{\"nombre\": \"Telefono\", \"cantidad\": 1, \"subtotal\": 150, \"productoId\": 3, \"precioUnitario\": 150}]'),(16,10,'2026-02-28 19:34:23',1,'[{\"nombre\": \"Camista\", \"cantidad\": 1, \"subtotal\": 10, \"productoId\": 2, \"precioUnitario\": 10}]'),(17,30,'2026-02-28 19:35:51',1,'[{\"nombre\": \"Camista\", \"cantidad\": 1, \"subtotal\": 10, \"productoId\": 2, \"precioUnitario\": 10}, {\"nombre\": \"Zapatos\", \"cantidad\": 1, \"subtotal\": 20, \"productoId\": 1, \"precioUnitario\": 20}]');
/*!40000 ALTER TABLE `orden` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2026-02-28 22:57:41
