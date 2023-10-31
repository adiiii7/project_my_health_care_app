# project_my_health_care_app

CREATE DATABASE healthcare;


USE healthcare;



CREATE TABLE User (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(20) NOT NULL
);


INSERT INTO User (username, password, role) VALUES
('Adi', 'adi@123', 'admin'),
('Rahul', 'ra@123', 'patient'),
('Siddhant', 'sid@123', 'patient');


CREATE TABLE Patient (
    patientid INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    phoneno VARCHAR(20) NOT NULL,
    age INT NOT NULL,
    gender VARCHAR(10) NOT NULL
);

INSERT INTO Patient (patientid, name, phoneno,age,gender) VALUES
(11,'Rahul',987,12,'M'),
(12,'Siddhant',786,14,'M'),
(13,'soni',127,20,'F');


CREATE TABLE Appointment (
    id INT AUTO_INCREMENT PRIMARY KEY,
    appointmentdate DATETIME NOT NULL,
    approvalstatus VARCHAR(20) NOT NULL,
    diagnosticTests TEXT,
    patient INT NOT NULL,
    diagnosticCenter INT NOT NULL,
    testResult INT
);


CREATE TABLE TestResult (
    id INT AUTO_INCREMENT PRIMARY KEY,
    testreading VARCHAR(255) NOT NULL,
    condition TEXT,
    appointment INT NOT NULL
);


CREATE TABLE DiagnosticTest (
    id INT AUTO_INCREMENT PRIMARY KEY,
    testname VARCHAR(255) NOT NULL,
    testprice DECIMAL(10, 2) NOT NULL,
    normalvalue VARCHAR(255) NOT NULL,
    units VARCHAR(20) NOT NULL,
    diagnosticCenter INT NOT NULL
);


CREATE TABLE DiagnosticCenter (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    contactno VARCHAR(20) NOT NULL,
    address TEXT,
    contactemail VARCHAR(255) NOT NULL,
    serviceoffered TEXT,
    tests TEXT
);
