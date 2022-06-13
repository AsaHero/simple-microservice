--
-- PostgreSQL database dump
--

-- Dumped from database version 14.3 (Ubuntu 14.3-0ubuntu0.22.04.1)
-- Dumped by pg_dump version 14.3 (Ubuntu 14.3-0ubuntu0.22.04.1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: uuid-ossp; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;


--
-- Name: EXTENSION "uuid-ossp"; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';


--
-- Name: layout; Type: TYPE; Schema: public; Owner: asad
--

CREATE TYPE public.layout AS ENUM (
    'QWERTY',
    'QWERTZ',
    'AZERTY'
);


ALTER TYPE public.layout OWNER TO asad;

--
-- Name: panel; Type: TYPE; Schema: public; Owner: asad
--

CREATE TYPE public.panel AS ENUM (
    'IPS',
    'OLED'
);


ALTER TYPE public.panel OWNER TO asad;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: cpu; Type: TABLE; Schema: public; Owner: asad
--

CREATE TABLE public.cpu (
    id integer NOT NULL,
    brand character varying NOT NULL,
    name character varying NOT NULL,
    number_cores integer NOT NULL,
    number_threads integer NOT NULL,
    min_ghz double precision NOT NULL,
    max_ghz double precision NOT NULL
);


ALTER TABLE public.cpu OWNER TO asad;

--
-- Name: cpu_id_seq; Type: SEQUENCE; Schema: public; Owner: asad
--

CREATE SEQUENCE public.cpu_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.cpu_id_seq OWNER TO asad;

--
-- Name: cpu_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: asad
--

ALTER SEQUENCE public.cpu_id_seq OWNED BY public.cpu.id;


--
-- Name: laptop; Type: TABLE; Schema: public; Owner: asad
--

CREATE TABLE public.laptop (
    id integer NOT NULL,
    laptop_id uuid DEFAULT public.uuid_generate_v4(),
    brand character varying NOT NULL,
    name character varying NOT NULL,
    cpu_id integer,
    ram json NOT NULL,
    gpu json NOT NULL,
    storage json NOT NULL,
    screen_id integer,
    keyboard json NOT NULL,
    weight json NOT NULL,
    price_usd double precision NOT NULL,
    release_year integer,
    update_at timestamp without time zone,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.laptop OWNER TO asad;

--
-- Name: laptop_id_seq; Type: SEQUENCE; Schema: public; Owner: asad
--

CREATE SEQUENCE public.laptop_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.laptop_id_seq OWNER TO asad;

--
-- Name: laptop_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: asad
--

ALTER SEQUENCE public.laptop_id_seq OWNED BY public.laptop.id;


--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: asad
--

CREATE TABLE public.schema_migrations (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


ALTER TABLE public.schema_migrations OWNER TO asad;

--
-- Name: screen; Type: TABLE; Schema: public; Owner: asad
--

CREATE TABLE public.screen (
    id integer NOT NULL,
    width integer NOT NULL,
    height integer NOT NULL,
    size_inch double precision NOT NULL,
    panel public.panel,
    multitouch boolean NOT NULL
);


ALTER TABLE public.screen OWNER TO asad;

--
-- Name: screen_id_seq; Type: SEQUENCE; Schema: public; Owner: asad
--

CREATE SEQUENCE public.screen_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.screen_id_seq OWNER TO asad;

--
-- Name: screen_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: asad
--

ALTER SEQUENCE public.screen_id_seq OWNED BY public.screen.id;


--
-- Name: cpu id; Type: DEFAULT; Schema: public; Owner: asad
--

ALTER TABLE ONLY public.cpu ALTER COLUMN id SET DEFAULT nextval('public.cpu_id_seq'::regclass);


--
-- Name: laptop id; Type: DEFAULT; Schema: public; Owner: asad
--

ALTER TABLE ONLY public.laptop ALTER COLUMN id SET DEFAULT nextval('public.laptop_id_seq'::regclass);


--
-- Name: screen id; Type: DEFAULT; Schema: public; Owner: asad
--

ALTER TABLE ONLY public.screen ALTER COLUMN id SET DEFAULT nextval('public.screen_id_seq'::regclass);


--
-- Data for Name: cpu; Type: TABLE DATA; Schema: public; Owner: asad
--

COPY public.cpu (id, brand, name, number_cores, number_threads, min_ghz, max_ghz) FROM stdin;
11	Intel	Core i5-9400F	5	10	2.3174369544198368	2.4926762522489536
12	Intel	Core i7-9750H	4	8	2.867845035572349	4.660719784741778
13	AMD	Ryzen 3 PRO 3200GE	8	16	2.176475664620252	4.516323502160212
14	Intel	Core i3-1005G1	2	4	3.1774743446847524	3.1876003097008985
15	Intel	Core i9-9980HK	2	4	2.298877649388697	3.64051409772269
16	AMD	Ryzen 7 PRO 2700U	3	6	3.4691132284571253	4.295463410242403
17	AMD	Ryzen 3 PRO 3200GE	3	6	2.0688590152412725	2.556288058054019
\.


--
-- Data for Name: laptop; Type: TABLE DATA; Schema: public; Owner: asad
--

COPY public.laptop (id, laptop_id, brand, name, cpu_id, ram, gpu, storage, screen_id, keyboard, weight, price_usd, release_year, update_at, created_at) FROM stdin;
1	ad3d1784-eb9b-4e4c-b209-0767ffe7cea3	Apple	Macbook Air	11	{"value":28,"unit":5}	[{"brand":"AMD","name":"RX 580","min_ghz":1.4365718762545245,"max_ghz":1.9444810342241852,"memory":{"value":4,"unit":5}}]	[{"driver":1,"memory":{"value":4,"unit":6}},{"driver":2,"memory":{"value":453,"unit":5}}]	3	{"layout":1,"backlit":true}	{"WeightKg":1.8649246690259447}	2945.2972971355775	2018	\N	2022-06-12 04:40:03.281843
2	cb3e1e11-8a58-4798-bbd1-c7b277526c2b	Apple	Macbook Air	12	{"value":18,"unit":5}	[{"brand":"AMD","name":"RX 590","min_ghz":1.08308878444662,"max_ghz":1.9513595912039516,"memory":{"value":5,"unit":5}}]	[{"driver":1,"memory":{"value":5,"unit":6}},{"driver":2,"memory":{"value":410,"unit":5}}]	4	{"layout":2,"backlit":true}	{"WeightKg":2.1759351525941386}	2861.849146238015	2017	\N	2022-06-12 13:48:38.215049
3	97682fc9-5fe3-4a84-b113-ef31c5722cc2	Dell	Alienware	13	{"value":33,"unit":5}	[{"brand":"AMD","name":"RX 580","min_ghz":1.3961194794459362,"max_ghz":1.6633193529791304,"memory":{"value":4,"unit":5}}]	[{"driver":1,"memory":{"value":5,"unit":6}},{"driver":2,"memory":{"value":622,"unit":5}}]	5	{"layout":3,"backlit":true}	{"WeightKg":1.1911904027988613}	2532.9034732868968	2015	\N	2022-06-13 10:37:15.083468
4	67d94443-27a6-4344-a015-4dd655534aac	Dell	Vostro	14	{"value":56,"unit":5}	[{"brand":"NVIDIA","name":"GTX 1070","min_ghz":1.3784638559014915,"max_ghz":1.6363021144518122,"memory":{"value":7,"unit":5}}]	[{"driver":1,"memory":{"value":2,"unit":6}},{"driver":2,"memory":{"value":737,"unit":5}}]	6	{"layout":3}	{"WeightKg":1.356755261318439}	2441.564240207498	2018	\N	2022-06-13 10:40:19.182855
5	b40ae1dc-063d-4686-a7cf-748364a9f1ed	Dell	Alienware	15	{"value":30,"unit":5}	[{"brand":"NVIDIA","name":"GTX 1660-Ti","min_ghz":1.216055865238543,"max_ghz":1.4448420667169093,"memory":{"value":8,"unit":5}}]	[{"driver":1,"memory":{"value":3,"unit":6}},{"driver":2,"memory":{"value":880,"unit":5}}]	7	{"layout":3,"backlit":true}	{"WeightKg":1.6751346287913067}	1889.403308358146	2019	\N	2022-06-13 10:47:54.045421
6	23c5f2ac-3596-4d9c-b882-adc013c92d67	HP	Thinkpad	16	{"value":41,"unit":5}	[{"brand":"NVIDIA","name":"GTX 1070","min_ghz":1.2591518904262027,"max_ghz":1.7471303474284887,"memory":{"value":3,"unit":5}}]	[{"driver":1,"memory":{"value":4,"unit":6}},{"driver":2,"memory":{"value":138,"unit":5}}]	8	{"layout":1}	{"WeightKg":2.913826278747165}	2368.830901606415	2020	\N	2022-06-13 10:59:37.332379
7	4a050c87-04a3-4293-b383-ae01d069791d	Dell	XPS	17	{"value":5,"unit":5}	[{"brand":"NVIDIA","name":"GTX 1070","min_ghz":1.2323558629807778,"max_ghz":1.6729957695764832,"memory":{"value":4,"unit":5}}]	[{"driver":1,"memory":{"value":5,"unit":6}},{"driver":2,"memory":{"value":733,"unit":5}}]	9	{"layout":3}	{"WeightKg":1.8217733062068153}	1960.9232563885332	2016	\N	2022-06-13 11:00:10.178799
\.


--
-- Data for Name: schema_migrations; Type: TABLE DATA; Schema: public; Owner: asad
--

COPY public.schema_migrations (version, dirty) FROM stdin;
1	f
\.


--
-- Data for Name: screen; Type: TABLE DATA; Schema: public; Owner: asad
--

COPY public.screen (id, width, height, size_inch, panel, multitouch) FROM stdin;
3	5640	3173	22.156042098999023	IPS	f
4	2467	1388	31.39840316772461	OLED	t
5	3224	1814	15.242265701293945	IPS	t
6	3408	1917	31.387147903442383	IPS	f
7	7191	4045	24.269268035888672	IPS	f
8	5699	3206	14.957919120788574	IPS	f
9	2332	1312	20.029401779174805	OLED	t
\.


--
-- Name: cpu_id_seq; Type: SEQUENCE SET; Schema: public; Owner: asad
--

SELECT pg_catalog.setval('public.cpu_id_seq', 17, true);


--
-- Name: laptop_id_seq; Type: SEQUENCE SET; Schema: public; Owner: asad
--

SELECT pg_catalog.setval('public.laptop_id_seq', 7, true);


--
-- Name: screen_id_seq; Type: SEQUENCE SET; Schema: public; Owner: asad
--

SELECT pg_catalog.setval('public.screen_id_seq', 9, true);


--
-- Name: cpu cpu_pkey; Type: CONSTRAINT; Schema: public; Owner: asad
--

ALTER TABLE ONLY public.cpu
    ADD CONSTRAINT cpu_pkey PRIMARY KEY (id);


--
-- Name: laptop laptop_pkey; Type: CONSTRAINT; Schema: public; Owner: asad
--

ALTER TABLE ONLY public.laptop
    ADD CONSTRAINT laptop_pkey PRIMARY KEY (id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: asad
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: screen screen_pkey; Type: CONSTRAINT; Schema: public; Owner: asad
--

ALTER TABLE ONLY public.screen
    ADD CONSTRAINT screen_pkey PRIMARY KEY (id);


--
-- Name: laptop laptop_cpu_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: asad
--

ALTER TABLE ONLY public.laptop
    ADD CONSTRAINT laptop_cpu_id_fkey FOREIGN KEY (cpu_id) REFERENCES public.cpu(id);


--
-- Name: laptop laptop_screen_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: asad
--

ALTER TABLE ONLY public.laptop
    ADD CONSTRAINT laptop_screen_id_fkey FOREIGN KEY (screen_id) REFERENCES public.screen(id);


--
-- PostgreSQL database dump complete
--

