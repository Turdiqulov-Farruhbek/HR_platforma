CREATE TABLE IF NOT EXISTS users(
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    phone_number VARCHAR(255) NOT NULL,
    birthday timestamp NOT NULL,
    gender VARCHAR(255) CHECK (gender IN ('male', 'female')) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0,
    UNIQUE (email, phone_number)
);


CREATE TABLE IF NOT EXISTS resume(
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    position VARCHAR(255) NOT NULL,                     --lavozimi
    experience INT DEFAULT 0,                           --tajribasi
    description TEXT NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0

);


CREATE TABLE IF NOT EXISTS company(
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    location VARCHAR(255) NOT NULL,
    workers INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE IF NOT EXISTS recruiter(
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    phone_number VARCHAR(255) NOT NULL,
    birthday timestamp NOT NULL,
    gender VARCHAR(255) CHECK (gender IN ('male', 'female')) NOT NULL,
    company_id UUID REFERENCES company(id),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE IF NOT EXISTS vacancy(
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    position VARCHAR(255) NOT NULL,                 --vakansiya qaysi yunalishda ekanligi
    min_exp INT DEFAULT 0,                          -- minimum tajriba
    company_id UUID REFERENCES company(id),
    description TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0

);


CREATE TABLE IF NOT EXISTS interview(
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    user_id UUID REFERENCES users(id),
    vacancy_id UUID REFERENCES vacancy(id),
    recruiter_id UUID REFERENCES recruiter(id),
    interview_date TIMESTAMP DEFAULT NOW(),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0

);

ALTER TABLE recruiter
ADD COLUMN created_at TIMESTAMP NOT NULL DEFAULT now();

ALTER TABLE recruiter
ADD COLUMN updated_at TIMESTAMP NOT NULL DEFAULT now();

ALTER TABLE recruiter 
ADD COLUMN deleted_at BIGINT NOT NULL DEFAULT 0;


INSERT INTO company (name, location, workers) VALUES 
('Ozbek Milliy Banki', 'Toshkent', 5000),
('Ozbek Energetika Kompaniyasi', 'Toshkent', 1200),
('Ozbek Neft va Gaz', 'Buxoro', 3000),
('Ozbek Temir Yollari', 'Samarqand', 1500),
('Ozbek Telekom', 'Navoiy', 800),
('Ozbek Qishloq Hojaligi', 'Fargona', 900),
('Ozbek Toqimachilik', 'Andijon', 700),
('Ozbek Kimyo Sanoati', 'Nukus', 1100),
('Ozbek Farmatsevtika', 'Qarshi', 400),
('Ozbek Turizm', 'Xiva', 300);


INSERT INTO recruiter (name, email, phone_number, birthday, gender, company_id) VALUES 
('Azizbek Karimov', 'azizbek.karimov@example.com', '+998901234567', '1985-05-15', 'male', '29fc303a-2f84-4af7-9033-b1a0a933c329'),
('Gulnora Ismoilova', 'gulnora.ismoilova@example.com', '+998902345678', '1990-07-20', 'female', '8f1d72fb-4d14-42cf-9fa9-8fc4e9c8367e'),
('Javlonbek Qodirov', 'javlonbek.qodirov@example.com', '+998903456789', '1982-03-10', 'male', 'b4ed2d35-9f8f-46b0-9710-263938be81a8'),
('Shahnoza Rahimova', 'shahnoza.rahimova@example.com', '+998904567890', '1988-11-30', 'female', '29fc303a-2f84-4af7-9033-b1a0a933c329'),
('Anvarbek Toshmatov', 'anvarbek.toshmatov@example.com', '+998905678901', '1979-01-25', 'male', '90ab9aa4-f00c-4def-a8ac-a1e6684f8a9a'),
('Feruza Abdurahmonova', 'feruza.abdurahmonova@example.com', '+998906789012', '1992-06-15', 'female', '29fc303a-2f84-4af7-9033-b1a0a933c329'),
('Rustambek Ergashev', 'rustambek.ergashev@example.com', '+998907890123', '1987-08-22', 'male', '3f62991f-8f11-457f-8015-772d32c88591'),
('Lola Qosimova', 'lola.qosimova@example.com', '+998908901234', '1991-02-05', 'female', '932f2e00-2cf5-4f4e-953e-e516996ce65c'),
('Dilshodbek Xudoyberdiyev', 'dilshodbek.xudoyberdiyev@example.com', '+998909012345', '1983-09-14', 'male', 'b4ed2d35-9f8f-46b0-9710-263938be81a8'),
('Oygul Murodova', 'oygul.murodova@example.com', '+998900123456', '1995-04-18', 'female', '90ab9aa4-f00c-4def-a8ac-a1e6684f8a9a');


INSERT INTO vacancy (name, position, min_exp, company_id, description) VALUES
('Dasturchi', 'Backend', 3, '3f62991f-8f11-457f-8015-772d32c88591', 'Java va Spring Framework bo''yicha 3 yillik tajriba talab qilinadi.'),
('Dizayner', 'Grafik', 2, 'b4ed2d35-9f8f-46b0-9710-263938be81a8', 'Photoshop va Illustrator bo''yicha tajribali dizayner izlanmoqda.'),
('Muhandis', 'Elektro', 5, '4180facb-9f26-4d68-951b-a0cbef84919f', 'Elektr tarmoqlari bo''yicha 5 yillik tajriba talab qilinadi.'),
('Hisobchi', 'Moliyaviy', 4, '8e937557-123d-450c-a9ed-9956676b0fd9', 'Moliya va hisob bo''yicha 4 yillik tajriba talab qilinadi.'),
('Boshqaruvchi', 'Loyiha', 6, '90ab9aa4-f00c-4def-a8ac-a1e6684f8a9a', 'Loyiha boshqaruvi bo''yicha 6 yillik tajriba talab qilinadi.'),
('Texnik', 'IT', 3, '85c929e1-776d-46b8-8e06-0711d8991edd', 'IT texnikasi bo''yicha 3 yillik tajriba talab qilinadi.'),
('Marketing', 'Raqamli', 4, '8f1d72fb-4d14-42cf-9fa9-8fc4e9c8367e', 'Raqamli marketing bo''yicha 4 yillik tajriba talab qilinadi.'),
('Oqituvchi', 'Ingliz tili', 2, '29fc303a-2f84-4af7-9033-b1a0a933c329', 'Ingliz tilini o''qitish bo''yicha 2 yillik tajriba talab qilinadi.'),
('Tibbiyot xodimi', 'Hamshira', 3, '405428e5-febc-412c-ae2b-becc48aa892f', 'Hamshiralik bo''yicha 3 yillik tajriba talab qilinadi.'),
('Savdo vakili', 'Savdo', 1, '90ab9aa4-f00c-4def-a8ac-a1e6684f8a9a', 'Savdo sohasida 1 yillik tajriba talab qilinadi.');


INSERT INTO resume (position, experience, description, user_id, created_at, updated_at, deleted_at) VALUES 
( 'Dasturchi', 5, 'Tajribali web dasturchi', '41539f89-cb37-4de1-8f0b-9340c6a9e931', NOW(), NOW(), 0),
( 'Dizayner', 3, 'Grafik dizayner', 'b6229219-b996-4b5e-91a6-2bd5db4f9f4f', NOW(), NOW(), 0),
( 'Loyiha menejeri', 8, 'Loyiha boshqaruvi', 'e157adae-daae-4e90-b9d4-a838f30bb5d2', NOW(), NOW(), 0),
( 'Tizim administratori', 6, 'Tarmoq va server boshqaruvi', '5888e3fa-8690-411d-a2a1-0ffcdd8758d0', NOW(), NOW(), 0),
( 'Malumotlar bazasi mutaxassisi', 4, 'SQL va malumotlar bazasi boshqaruvi', 'd4ccd783-cf3c-417a-914c-dc550388352e', NOW(), NOW(), 0),
( 'Mobil ilova dasturchisi', 2, 'Android va iOS dasturchisi', '12fe199b-1180-4280-bf1a-0416cd006a85', NOW(), NOW(), 0),
( 'SEO mutaxassisi', 3, 'SEO va internet marketing', '1fe46ae0-9866-4335-bbbb-12e0898a58ed', NOW(), NOW(), 0),
( 'Texnik yozuvchi', 5, 'Texnik hujjatlar tayyorlash', 'fa0b639f-a00d-4ca6-a500-5e2893c8b13d', NOW(), NOW(), 0),
( 'Kiberxavfsizlik mutaxassisi', 6, 'Kiberxavfsizlik va malumotlarni himoya qilish', 'e1a53365-7d7b-45e1-bc02-d4ead9b79ac8', NOW(), NOW(), 0),
( 'DevOps injeneri', 4, 'DevOps va CI/CD jarayonlari', '89ab257e-a8a4-4d07-a162-6efa8676a475', NOW(), NOW(), 0);



INSERT INTO users (name, email, phone_number, birthday, gender, created_at, updated_at, deleted_at) VALUES 
('Olimjon', 'olimjon@gmail.com', '+998901234567', '1990-01-01', 'male', NOW(), NOW(), 0),
('Zarifa', 'zarifa@gmail.com', '+998901234568', '1992-02-02', 'female', NOW(), NOW(), 0),
('Farhod', 'farhod@gmail.com', '+998901234569', '1988-03-03', 'male', NOW(), NOW(), 0),
('Nigora', 'nigora@gmail.com', '+998901234570', '1995-04-04', 'female', NOW(), NOW(), 0),
('Jasur', 'jasur@gmail.com', '+998901234571', '1993-05-05', 'male', NOW(), NOW(), 0),
('Gulnora', 'gulnora@gmail.com', '+998901234572', '1989-06-06', 'female', NOW(), NOW(), 0),
('Aziz', 'aziz@gmail.com', '+998901234573', '1991-07-07', 'male', NOW(), NOW(), 0),
('Shahnoza', 'shahnoza@gmail.com', '+998901234574', '1994-08-08', 'female', NOW(), NOW(), 0),
('Rustam', 'rustam@gmail.com', '+998901234575', '1990-09-09', 'male', NOW(), NOW(), 0),
('Madina', 'madina@gmail.com', '+998901234576', '1992-10-10', 'female', NOW(), NOW(), 0);


INSERT INTO resume (position, experience, description, user_id) VALUES 
('Dasturchi', 5, 'Web va mobil ilovalar ishlab chiqish', '89ab257e-a8a4-4d07-a162-6efa8676a475'),
('Muhandis', 8, 'Mexanik tizimlarni loyihalash va ishlab chiqarish', 'e1a53365-7d7b-45e1-bc02-d4ead9b79ac8'),
('Loyiha menejeri', 10, 'Loyihalarni boshqarish va nazorat qilish', 'fa0b639f-a00d-4ca6-a500-5e2893c8b13d'),
('Marketing mutaxassisi', 6, 'Reklama va marketing strategiyalarini ishlab chiqish', '1fe46ae0-9866-4335-bbbb-12e0898a58ed'),
('Hisobchi', 7, 'Moliyaviy hisobotlarni tayyorlash', '12fe199b-1180-4280-bf1a-0416cd006a85'),
('Dizayner', 4, 'Grafik va veb dizayn', 'd4ccd783-cf3c-417a-914c-dc550388352e'),
('O‘qituvchi', 12, 'Matematika va fizika fanlarini o‘qitish', '5888e3fa-8690-411d-a2a1-0ffcdd8758d0'),
('Tibbiyot xodimi', 15, 'Bemorlarni davolash va parvarish qilish', 'e157adae-daae-4e90-b9d4-a838f30bb5d2'),
('Huquqshunos', 9, 'Huquqiy maslahat va xizmatlar ko‘rsatish', '41539f89-cb37-4de1-8f0b-9340c6a9e931'),
('Tarjimon', 3, 'Hujjatlar va matnlarni tarjima qilish', 'b6229219-b996-4b5e-91a6-2bd5db4f9f4f');

SELECT v.id, v.name, v.position, v.min_exp, i.id, i.user_id, i.recruiter_id, i.interview_date
FROM vacancy v
JOIN interview i ON i.vacancy_id = v.id
WHERE v.company_id = '405428e5-febc-412c-ae2b-becc48aa892f';