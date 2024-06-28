INSERT INTO wheather(city, country, local_time, tempC, windKmph, humidity)
VALUES
    ('Xonqa', 'Uzbekiston', '2024-06-27', 43.6, 10, 12),
    ('Tashkent', 'Uzbekiston', '2024-06-27', 40.6, 12, 22),
    ('Samarqand', 'Uzbekiston', '2024-06-27', 38.6, 17, 27);

insert into transports(bus_number, stations)
values
    (114, array['Shuhrat dokoni', 'Boshliq mavzesi']),
    (69, array['Shuhrat dokoni', 'Boshliq mavzesi']),
    (94, array['Aeraport', 'Najot talim oldi']);
