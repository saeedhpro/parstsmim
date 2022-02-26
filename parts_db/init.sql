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


CREATE DATABASE IF NOT EXISTS parts DEFAULT CHARSET 'utf8mb4' COLLATE 'utf8mb4_unicode_ci';

CREATE USER 'user'@'%' IDENTIFIED BY PASSWORD 'password';
GRANT ALL PRIVILEGES ON parts.* TO 'user'@'%' WITH GRANT OPTION;
GRANT ALL PRIVILEGES ON parts.* TO 'user'@'%localhost' WITH GRANT OPTION;

FLUSH PRIVILEGES;

USE parts;


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `parts`
--

-- --------------------------------------------------------

--
-- Table structure for table `parts`
--

CREATE TABLE `parts` (
                         `id` bigint UNSIGNED NOT NULL,
                         `automobile_id` bigint UNSIGNED NOT NULL,
                         `name` varchar(128) COLLATE utf8mb4_persian_ci DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_persian_ci;

--
-- Dumping data for table `parts`
--

INSERT INTO `parts` (`id`, `automobile_id`, `name`) VALUES
                                                        (1, 1, 'part 1'),
                                                        (2, 1, 'part 2'),
                                                        (3, 1, 'part 3'),
                                                        (4, 1, 'part 4'),
                                                        (5, 1, 'part 5'),
                                                        (6, 1, 'part 6');

-- --------------------------------------------------------

--
-- Table structure for table `part_files`
--

CREATE TABLE `part_files` (
                              `id` bigint UNSIGNED NOT NULL,
                              `part_id` bigint UNSIGNED NOT NULL,
                              `name` varchar(128) COLLATE utf8mb4_persian_ci NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_persian_ci;

--
-- Dumping data for table `part_files`
--

INSERT INTO `part_files` (`id`, `part_id`, `name`) VALUES
                                                       (1, 1, 'part_1_file_1.png'),
                                                       (2, 1, 'part_1_file_2.png'),
                                                       (3, 1, 'part_1_file_3.png'),
                                                       (4, 2, 'part_2_file_1.png'),
                                                       (5, 3, 'part_3_file_1.png'),
                                                       (6, 4, 'part_4_file_1.png'),
                                                       (7, 4, 'part_4_file_2.png');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `parts`
--
ALTER TABLE `parts`
    ADD PRIMARY KEY (`id`);

--
-- Indexes for table `part_files`
--
ALTER TABLE `part_files`
    ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `parts`
--
ALTER TABLE `parts`
    MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT for table `part_files`
--
ALTER TABLE `part_files`
    MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;