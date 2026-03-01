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
-- Table structure for table `producto`
--

DROP TABLE IF EXISTS `producto`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `producto` (
  `ProductoID` int NOT NULL AUTO_INCREMENT,
  `Nombre` varchar(45) NOT NULL,
  `Descripcion` varchar(100) NOT NULL,
  `Precio` float NOT NULL,
  `Stock` int NOT NULL,
  `CategoriaID` int NOT NULL,
  PRIMARY KEY (`ProductoID`),
  KEY `fk_CategoríaID_idx` (`CategoriaID`),
  CONSTRAINT `fk_CategoríaID` FOREIGN KEY (`CategoriaID`) REFERENCES `categoria` (`CategoriaID`)
) ENGINE=InnoDB AUTO_INCREMENT=104 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `producto`
--

LOCK TABLES `producto` WRITE;
/*!40000 ALTER TABLE `producto` DISABLE KEYS */;
INSERT INTO `producto` VALUES (1,'Zapatos','Zapatos',20,50,2),(2,'Camista','Camistea',10,20,2),(3,'Telefono','Telefono',150,25,1),(54,'Laptop HP 15\"','Laptop HP con 8GB RAM y 512GB SSD',750,15,1),(55,'Smartphone Samsung A54','Telefono inteligente 128GB',320,25,1),(56,'Audifonos Bluetooth Sony','Audifonos inalambricos con cancelacion de ruido',120,40,1),(57,'Smart TV 50\" LG','Televisor 4K UHD con HDR',600,10,1),(58,'Tablet Lenovo M10','Tablet 10 pulgadas 64GB',210,18,1),(59,'Mouse Logitech','Mouse inalambrico ergonomico',25,50,1),(60,'Teclado Mecanico Redragon','Teclado gamer retroiluminado',70,30,1),(61,'Camara Web HD','Camara 1080p para videollamadas',45,35,1),(62,'Disco Duro Externo 1TB','Almacenamiento portatil USB 3.0',80,20,1),(63,'Power Bank 20000mAh','Bateria externa carga rapida',35,45,1),(64,'Camiseta Basica Blanca','Camiseta 100% algodon',15,60,2),(65,'Jeans Slim Fit','Pantalon jean azul oscuro',40,35,2),(66,'Chaqueta Deportiva','Chaqueta ligera impermeable',55,20,2),(67,'Vestido Casual','Vestido comodo para uso diario',30,25,2),(68,'Sudadera con Capucha','Sudadera unisex',35,30,2),(69,'Camisa Formal','Camisa manga larga elegante',28,22,2),(70,'Short Deportivo','Short transpirable para ejercicio',18,40,2),(71,'Blusa Floral','Blusa ligera estampada',22,27,2),(72,'Pijama Algodon','Pijama comoda para dormir',25,33,2),(73,'Abrigo Invierno','Abrigo grueso termico',85,12,2),(74,'Zapatillas Running Nike','Calzado deportivo para correr',95,20,3),(75,'Botas Cuero','Botas resistentes de cuero',120,15,3),(76,'Sandalias Playa','Sandalias ligeras y comodas',20,50,3),(77,'Zapatos Formales','Zapatos elegantes para oficina',75,18,3),(78,'Tenis Casual','Tenis urbanos comodos',60,25,3),(79,'Tacones Altos','Tacones elegantes 8cm',55,14,3),(80,'Botines Mujer','Botines modernos de moda',80,16,3),(81,'Crocs Clasicos','Calzado ligero y resistente al agua',35,28,3),(82,'Zapatillas Skate','Zapatillas resistentes para skate',65,19,3),(83,'Pantuflas Hogar','Pantuflas suaves y comodas',18,40,3),(84,'Reloj Deportivo','Reloj resistente al agua',50,30,4),(85,'Gorra Ajustable','Gorra casual unisex',15,45,4),(86,'Mochila Escolar','Mochila con multiples compartimentos',35,20,4),(87,'Cinturon Cuero','Cinturon clasico negro',20,38,4),(88,'Lentes de Sol','Proteccion UV400',25,50,4),(89,'Bolso de Mano','Bolso elegante para dama',45,18,4),(90,'Pulsera Acero','Pulsera resistente y moderna',12,60,4),(91,'Collar Plateado','Collar fino con dije',18,42,4),(92,'Billetera Hombre','Billetera cuero genuino',30,26,4),(93,'Bufanda Invierno','Bufanda tejida abrigada',22,29,4),(94,'Licuadora Oster','Licuadora 3 velocidades',65,15,5),(95,'Microondas 20L','Microondas digital 700W',110,10,5),(96,'Juego de Sartenes','Set antiadherente 3 piezas',75,18,5),(97,'Lampara LED','Lampara de escritorio regulable',28,35,5),(98,'Alfombra Sala','Alfombra decorativa 2x3m',90,12,5),(99,'Juego de Sabana','Sabana matrimonial 100% algodon',40,22,5),(100,'Cafetera Electrica','Cafetera programable 12 tazas',85,14,5),(101,'Ventilador Torre','Ventilador silencioso 3 velocidades',95,9,5),(102,'Organizador Plastico','Caja organizadora multiuso',15,50,5),(103,'Espejo Decorativo','Espejo moderno para pared',55,17,5);
/*!40000 ALTER TABLE `producto` ENABLE KEYS */;
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
