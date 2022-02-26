-- phpMyAdmin SQL Dump
-- version 4.9.5deb2
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Generation Time: Feb 14, 2022 at 07:05 AM
-- Server version: 8.0.28-0ubuntu0.20.04.3
-- PHP Version: 7.4.3

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


CREATE DATABASE IF NOT EXISTS automobiles DEFAULT CHARSET 'utf8mb4' COLLATE 'utf8mb4_unicode_ci';

CREATE USER 'user'@'%' IDENTIFIED BY PASSWORD 'password';
GRANT ALL PRIVILEGES ON parts.* TO 'user'@'%' WITH GRANT OPTION;
GRANT ALL PRIVILEGES ON parts.* TO 'user'@'%localhost' WITH GRANT OPTION;

FLUSH PRIVILEGES;

USE automobiles;


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `automobiles`
--

-- --------------------------------------------------------

--
-- Table structure for table `automobiles`
--

CREATE TABLE `automobiles` (
                         `id` bigint UNSIGNED NOT NULL,
                         `manufacture` varchar(128) COLLATE utf8mb4_persian_ci DEFAULT NULL,
                         `model` bigint UNSIGNED NOT NULL,
                         `type` varchar(128) COLLATE utf8mb4_persian_ci DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_persian_ci;

--
-- Dumping data for table `parts`
--

INSERT INTO `automobiles` (`id`, `manufacture`, `model`, `type`) VALUES (1, 'manufacture', 1374, 'type 1'), (2, 'manufacture', 1384, 'type 2'),(3, 'manufacture', 1392, 'type 3');

-- --------------------------------------------------------

--
-- Indexes for dumped tables
--

--
-- Indexes for table `automobiles`
--
ALTER TABLE `automobiles`
    ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `automobiles`
--
ALTER TABLE `automobiles`
    MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;