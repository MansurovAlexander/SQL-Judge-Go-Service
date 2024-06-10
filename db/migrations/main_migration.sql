--
-- PostgreSQL database dump
--

-- Dumped from database version 14.11
-- Dumped by pg_dump version 14.5

DROP DATABASE IF EXISTS sqljudge;
CREATE DATABASE sqljudge;
\c sqljudge

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: assign; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.assign (
    id bigint NOT NULL,
    assign_id bigint NOT NULL,
    time_limit integer,
    memory_limit integer,
    correct_script text NOT NULL,
    db_id bigint NOT NULL,
    subtask_id numeric(3,0) DEFAULT '-1'::integer NOT NULL
);


--
-- Name: assign_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.assign_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: assign_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.assign_id_seq OWNED BY public.assign.id;


--
-- Name: banned_words_to_assign; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.banned_words_to_assign (
    id bigint NOT NULL,
    assign_id bigint NOT NULL,
    subtask_id integer NOT NULL,
    banned_words text DEFAULT ''::text,
    admission_words text DEFAULT ''::text
);


--
-- Name: banned_words_to_assign_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.banned_words_to_assign_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: banned_words_to_assign_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.banned_words_to_assign_id_seq OWNED BY public.banned_words_to_assign.id;


--
-- Name: databases; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.databases (
    id bigint NOT NULL,
    name character varying(50) DEFAULT 'Not restored'::character varying,
    description character varying(50) DEFAULT NULL::character varying,
    dbms_id integer NOT NULL,
    file_name character varying(255)
);


--
-- Name: databases_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.databases_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: databases_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.databases_id_seq OWNED BY public.databases.id;


--
-- Name: dbms; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.dbms (
    id integer NOT NULL,
    name character varying(10) NOT NULL
);


--
-- Name: dbms_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.dbms_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: dbms_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.dbms_id_seq OWNED BY public.dbms.id;


--
-- Name: status; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.status (
    id integer NOT NULL,
    name character varying(255) NOT NULL
);


--
-- Name: status_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.status_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: status_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.status_id_seq OWNED BY public.status.id;


--
-- Name: submission; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.submission (
    id bigint NOT NULL,
    student_id bigint NOT NULL,
    "time" bigint,
    memory bigint,
    script text NOT NULL,
    assign_id bigint NOT NULL,
    submission_id bigint,
    status_id integer NOT NULL,
    subtask_id numeric(3,0) NOT NULL
);


--
-- Name: submission_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.submission_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: submission_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.submission_id_seq OWNED BY public.submission.id;


--
-- Name: assign id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.assign ALTER COLUMN id SET DEFAULT nextval('public.assign_id_seq'::regclass);


--
-- Name: banned_words_to_assign id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.banned_words_to_assign ALTER COLUMN id SET DEFAULT nextval('public.banned_words_to_assign_id_seq'::regclass);


--
-- Name: databases id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.databases ALTER COLUMN id SET DEFAULT nextval('public.databases_id_seq'::regclass);


--
-- Name: dbms id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.dbms ALTER COLUMN id SET DEFAULT nextval('public.dbms_id_seq'::regclass);


--
-- Name: status id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.status ALTER COLUMN id SET DEFAULT nextval('public.status_id_seq'::regclass);


--
-- Name: submission id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.submission ALTER COLUMN id SET DEFAULT nextval('public.submission_id_seq'::regclass);


--
-- Data for Name: assign; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.assign (id, assign_id, time_limit, memory_limit, correct_script, db_id, subtask_id) FROM stdin;
\.


--
-- Data for Name: banned_words_to_assign; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.banned_words_to_assign (id, assign_id, subtask_id, banned_words, admission_words) FROM stdin;
\.


--
-- Data for Name: databases; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.databases (id, name, description, dbms_id, file_name) FROM stdin;
\.


--
-- Data for Name: dbms; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.dbms (id, name) FROM stdin;
1	postgres
2	mysql
\.


--
-- Data for Name: status; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.status (id, name) FROM stdin;
0	Bad word
1	Admission word
2	Correct answer
3	Wrong answer
4	Time limit exceeded
5	Memory limit exceeded
6	Unknown error
7	Not checked
8	Checked
\.


--
-- Data for Name: submission; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.submission (id, student_id, "time", memory, script, assign_id, submission_id, status_id, subtask_id) FROM stdin;
\.


--
-- Name: assign_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.assign_id_seq', 68, true);


--
-- Name: banned_words_to_assign_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.banned_words_to_assign_id_seq', 1, true);


--
-- Name: databases_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.databases_id_seq', 45, true);


--
-- Name: dbms_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.dbms_id_seq', 2, true);


--
-- Name: status_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.status_id_seq', 4, true);


--
-- Name: submission_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.submission_id_seq', 63, true);


--
-- Name: assign assign_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.assign
    ADD CONSTRAINT assign_pkey PRIMARY KEY (id);


--
-- Name: banned_words_to_assign banned_words_to_assign_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.banned_words_to_assign
    ADD CONSTRAINT banned_words_to_assign_pkey PRIMARY KEY (id);


--
-- Name: databases databases_file_name_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.databases
    ADD CONSTRAINT databases_file_name_key UNIQUE (file_name);


--
-- Name: databases databases_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.databases
    ADD CONSTRAINT databases_pkey PRIMARY KEY (id);


--
-- Name: dbms dbms_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.dbms
    ADD CONSTRAINT dbms_pkey PRIMARY KEY (id);


--
-- Name: status status_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.status
    ADD CONSTRAINT status_pkey PRIMARY KEY (id);


--
-- Name: submission submission_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.submission
    ADD CONSTRAINT submission_pkey PRIMARY KEY (id);


--
-- Name: assign db_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.assign
    ADD CONSTRAINT db_id_fk FOREIGN KEY (db_id) REFERENCES public.databases(id) ON DELETE CASCADE;


--
-- Name: databases dbms_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.databases
    ADD CONSTRAINT dbms_id_fk FOREIGN KEY (dbms_id) REFERENCES public.dbms(id) ON DELETE CASCADE;


--
-- Name: submission submission_status_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.submission
    ADD CONSTRAINT submission_status_id_fkey FOREIGN KEY (status_id) REFERENCES public.status(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

