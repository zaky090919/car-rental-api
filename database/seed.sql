-- Customers
INSERT INTO customers (name, nik, phone_number) VALUES
('Wawan Hermawan','3372093912739','081237123682'),
('Philip Walker','3372093912785','081237123683'),
('Hugo Fleming','3372093912800','081237123684'),
('Maximillian Mendez','3372093912848','081237123685'),
('Felix Dixon','3372093912851','081237123686'),
('Nicholas Riddle','3372093912929','081237123687'),
('Stephen Wheeler','3372093912976','081237123688'),
('Roy Brennan','3372093913022','081237123689'),
('Eliza Le','3372093913106','081237123690'),
('Jesse Taylor','3372093913126','081237123691'),
('Damien Kaufman','3372093913202','081237123692'),
('Ayesha Richardson','3372093913257','081237123693'),
('Margaret Stokes','3372093913262','081237123694'),
('Sara Livingston','3372093913268','081237123695'),
('Callie Townsend','3372093913281','081237123696'),
('Lilly Fischer','3372093913325','081237123697'),
('Theresa Barton','3372093913335','081237123698'),
('Mia Curtis','3372093913343','081237123699'),
('Flora Barlow','3372093913400','081237123700'),
('Vanessa Patton','3372093913434','081237123701');

-- Cars
INSERT INTO cars (brand, model, year, price_per_day, status) VALUES
('Toyota','Avanza',2022,300000,'available'),
('Toyota','Innova',2023,450000,'available'),
('Honda','Brio',2021,250000,'available'),
('Honda','HRV',2022,400000,'available'),
('Suzuki','Ertiga',2023,320000,'available'),
('Suzuki','XL7',2022,380000,'available'),
('Mitsubishi','Xpander',2023,350000,'available'),
('Mitsubishi','Pajero',2021,800000,'available'),
('Daihatsu','Sigra',2022,220000,'available'),
('Daihatsu','Terios',2023,330000,'available');

-- Bookings
INSERT INTO bookings (customer_id, car_id, start_date, end_date, total_price) VALUES
(1,1,'2026-03-20','2026-03-22',0),
(2,3,'2026-03-21','2026-03-23',0),
(3,5,'2026-03-22','2026-03-25',0),
(4,2,'2026-03-24','2026-03-27',0),
(5,7,'2026-03-25','2026-03-28',0),
(6,4,'2026-03-26','2026-03-29',0),
(7,6,'2026-03-27','2026-03-30',0),
(8,8,'2026-03-28','2026-03-31',0),
(9,9,'2026-03-29','2026-04-01',0),
(10,10,'2026-03-30','2026-04-02',0);