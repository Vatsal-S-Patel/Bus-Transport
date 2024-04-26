--
-- PostgreSQL database dump
--

-- Dumped from database version 14.11 (Ubuntu 14.11-0ubuntu0.22.04.1)
-- Dumped by pg_dump version 14.11 (Ubuntu 14.11-0ubuntu0.22.04.1)

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
-- Name: arrays; Type: SCHEMA; Schema: -; Owner: vatsal
--

CREATE SCHEMA arrays;


ALTER SCHEMA arrays OWNER TO vatsal;

--
-- Name: data_definition; Type: SCHEMA; Schema: -; Owner: vatsal
--

CREATE SCHEMA data_definition;


ALTER SCHEMA data_definition OWNER TO vatsal;

--
-- Name: doc5; Type: SCHEMA; Schema: -; Owner: vatsal
--

CREATE SCHEMA doc5;


ALTER SCHEMA doc5 OWNER TO vatsal;

--
-- Name: gormbasics; Type: SCHEMA; Schema: -; Owner: vatsal
--

CREATE SCHEMA gormbasics;


ALTER SCHEMA gormbasics OWNER TO vatsal;

--
-- Name: gormbasics1; Type: SCHEMA; Schema: -; Owner: vatsal
--

CREATE SCHEMA gormbasics1;


ALTER SCHEMA gormbasics1 OWNER TO vatsal;

--
-- Name: gormbasics2; Type: SCHEMA; Schema: -; Owner: vatsal
--

CREATE SCHEMA gormbasics2;


ALTER SCHEMA gormbasics2 OWNER TO vatsal;

--
-- Name: gormschool; Type: SCHEMA; Schema: -; Owner: vatsal
--

CREATE SCHEMA gormschool;


ALTER SCHEMA gormschool OWNER TO vatsal;

--
-- Name: gormschoolproject; Type: SCHEMA; Schema: -; Owner: vatsal
--

CREATE SCHEMA gormschoolproject;


ALTER SCHEMA gormschoolproject OWNER TO vatsal;

--
-- Name: httpnet; Type: SCHEMA; Schema: -; Owner: vatsal
--

CREATE SCHEMA httpnet;


ALTER SCHEMA httpnet OWNER TO vatsal;

--
-- Name: products; Type: SCHEMA; Schema: -; Owner: vatsal
--

CREATE SCHEMA products;


ALTER SCHEMA products OWNER TO vatsal;

--
-- Name: school; Type: SCHEMA; Schema: -; Owner: vatsal
--

CREATE SCHEMA school;


ALTER SCHEMA school OWNER TO vatsal;

--
-- Name: sql; Type: SCHEMA; Schema: -; Owner: vatsal
--

CREATE SCHEMA sql;


ALTER SCHEMA sql OWNER TO vatsal;

--
-- Name: task6muxgorm; Type: SCHEMA; Schema: -; Owner: vatsal
--

CREATE SCHEMA task6muxgorm;


ALTER SCHEMA task6muxgorm OWNER TO vatsal;

--
-- Name: todo_cli; Type: SCHEMA; Schema: -; Owner: vatsal
--

CREATE SCHEMA todo_cli;


ALTER SCHEMA todo_cli OWNER TO vatsal;

--
-- Name: transport; Type: SCHEMA; Schema: -; Owner: vatsal
--

CREATE SCHEMA transport;


ALTER SCHEMA transport OWNER TO vatsal;

--
-- Name: calculate_age(); Type: FUNCTION; Schema: public; Owner: vatsal
--

CREATE FUNCTION public.calculate_age() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    NEW.age := DATE_PART('year', CURRENT_DATE) - DATE_PART('year', NEW.birth_date);
    RETURN NEW;
END;
$$;


ALTER FUNCTION public.calculate_age() OWNER TO vatsal;

--
-- Name: calculate_sum(integer, integer); Type: FUNCTION; Schema: public; Owner: vatsal
--

CREATE FUNCTION public.calculate_sum(a integer, b integer) RETURNS integer
    LANGUAGE plpgsql
    AS $$
BEGIN
	RETURN a + b;
END;
$$;


ALTER FUNCTION public.calculate_sum(a integer, b integer) OWNER TO vatsal;

--
-- Name: funcforroute(); Type: FUNCTION; Schema: public; Owner: vatsal
--

CREATE FUNCTION public.funcforroute() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    -- Drop the existing view if it exists
    EXECUTE 'DROP VIEW IF EXISTS bustransportsystem';

    -- Recreate the view with the updated data
    EXECUTE '
        CREATE VIEW bustransportsystem AS
        SELECT bus_id, v2.route_id as route_id, station_id, v2.name as route_name, 
               departure_time, source, destination, status, station_order, l.lat, 
               l.long, l.name as station_name  
        FROM (
            SELECT s.bus_id, v1.route_id, v1.name, s.departure_time, v1."source", 
                   v1.destination, v1.status, v1.station_order, station_id 
            FROM (
                SELECT r.id as route_id, r.name, rs.station_id, r.source, r.destination, 
                       r.status, station_order 
                FROM transport.routestations AS rs 
                INNER JOIN transport.route as r ON rs.route_id = r.id
            ) as v1 
            INNER JOIN transport.schedule as s ON s.route_id = v1.route_id
        ) as v2 
        INNER JOIN transport.station as l ON v2.station_id = l.id';

    RETURN NEW;
END;
$$;


ALTER FUNCTION public.funcforroute() OWNER TO vatsal;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: student; Type: TABLE; Schema: school; Owner: vatsal
--

CREATE TABLE school.student (
    id integer NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    name character varying(255),
    address character varying(40),
    city character varying(40),
    pincode character varying(6),
    birth_date date
);


ALTER TABLE school.student OWNER TO vatsal;

--
-- Name: get_student_info(); Type: FUNCTION; Schema: public; Owner: vatsal
--

CREATE FUNCTION public.get_student_info() RETURNS SETOF school.student
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY SELECT *FROM school.student WHERE school.student.id = 1;
END;
$$;


ALTER FUNCTION public.get_student_info() OWNER TO vatsal;

--
-- Name: get_student_info(integer); Type: FUNCTION; Schema: public; Owner: vatsal
--

CREATE FUNCTION public.get_student_info(id1 integer) RETURNS TABLE(id integer, name text, age integer)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY SELECT id, name, age FROM school.student WHERE id = id1;
END;
$$;


ALTER FUNCTION public.get_student_info(id1 integer) OWNER TO vatsal;

--
-- Name: print_numbers(); Type: FUNCTION; Schema: public; Owner: vatsal
--

CREATE FUNCTION public.print_numbers() RETURNS void
    LANGUAGE plpgsql
    AS $$
DECLARE
    i INTEGER := 1;
BEGIN
    LOOP
        EXIT WHEN i > 5;
        RAISE NOTICE '%', i;
        i := i + 1;
    END LOOP;
END;
$$;


ALTER FUNCTION public.print_numbers() OWNER TO vatsal;

--
-- Name: students(); Type: FUNCTION; Schema: public; Owner: vatsal
--

CREATE FUNCTION public.students() RETURNS SETOF school.student
    LANGUAGE plpgsql
    AS $$
BEGIN
	RETURN QUERY SELECT * FROM school.student;
END;
$$;


ALTER FUNCTION public.students() OWNER TO vatsal;

--
-- Name: students(integer); Type: FUNCTION; Schema: public; Owner: vatsal
--

CREATE FUNCTION public.students(id1 integer) RETURNS text
    LANGUAGE plpgsql
    AS $$
DECLARE
    student_name TEXT;
BEGIN
	 SELECT name INTO student_name FROM school.student WHERE school.student.id = id1;
   RETURN student_name;
END;
$$;


ALTER FUNCTION public.students(id1 integer) OWNER TO vatsal;

--
-- Name: updateviewroute(); Type: FUNCTION; Schema: public; Owner: vatsal
--

CREATE FUNCTION public.updateviewroute() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    -- Drop the existing view if it exists
    EXECUTE 'DROP VIEW IF EXISTS bustransportsystem';

    -- Recreate the view with the updated data
    EXECUTE '
        CREATE VIEW bustransportsystem AS
        SELECT bus_id, v2.route_id as route_id, station_id, v2.name as route_name, 
               departure_time, source, destination, status, station_order, l.lat, 
               l.long, l.name as station_name  
        FROM (
            SELECT s.bus_id, v1.route_id, v1.name, s.departure_time, v1."source", 
                   v1.destination, v1.status, v1.station_order, station_id 
            FROM (
                SELECT r.id as route_id, r.name, rs.station_id, r.source, r.destination, 
                       r.status, station_order 
                FROM transport.routestations AS rs 
                INNER JOIN transport.route as r ON rs.route_id = r.id
            ) as v1 
            INNER JOIN transport.schedule as s ON s.route_id = v1.route_id
        ) as v2 
        INNER JOIN transport.station as l ON v2.station_id = l.id';

    RETURN NEW;
END;
$$;


ALTER FUNCTION public.updateviewroute() OWNER TO vatsal;

--
-- Name: schools; Type: TABLE; Schema: arrays; Owner: vatsal
--

CREATE TABLE arrays.schools (
    id integer NOT NULL,
    names text[],
    ranks integer[]
);


ALTER TABLE arrays.schools OWNER TO vatsal;

--
-- Name: schools_id_seq; Type: SEQUENCE; Schema: arrays; Owner: vatsal
--

CREATE SEQUENCE arrays.schools_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE arrays.schools_id_seq OWNER TO vatsal;

--
-- Name: schools_id_seq; Type: SEQUENCE OWNED BY; Schema: arrays; Owner: vatsal
--

ALTER SEQUENCE arrays.schools_id_seq OWNED BY arrays.schools.id;


--
-- Name: generate_column; Type: TABLE; Schema: data_definition; Owner: vatsal
--

CREATE TABLE data_definition.generate_column (
    data_km integer,
    data_meter integer GENERATED ALWAYS AS ((data_km * 1000)) STORED
);


ALTER TABLE data_definition.generate_column OWNER TO vatsal;

--
-- Name: deestinct; Type: TABLE; Schema: doc5; Owner: vatsal
--

CREATE TABLE doc5.deestinct (
    id integer NOT NULL,
    name character varying(80)
);


ALTER TABLE doc5.deestinct OWNER TO vatsal;

--
-- Name: students; Type: TABLE; Schema: gormbasics; Owner: vatsal
--

CREATE TABLE gormbasics.students (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text,
    class bigint DEFAULT 18,
    age bigint DEFAULT 5
);


ALTER TABLE gormbasics.students OWNER TO vatsal;

--
-- Name: students_id_seq; Type: SEQUENCE; Schema: gormbasics; Owner: vatsal
--

CREATE SEQUENCE gormbasics.students_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE gormbasics.students_id_seq OWNER TO vatsal;

--
-- Name: students_id_seq; Type: SEQUENCE OWNED BY; Schema: gormbasics; Owner: vatsal
--

ALTER SEQUENCE gormbasics.students_id_seq OWNED BY gormbasics.students.id;


--
-- Name: category; Type: TABLE; Schema: gormbasics1; Owner: vatsal
--

CREATE TABLE gormbasics1.category (
    name text NOT NULL,
    remark text
);


ALTER TABLE gormbasics1.category OWNER TO vatsal;

--
-- Name: creditcard; Type: TABLE; Schema: gormbasics1; Owner: vatsal
--

CREATE TABLE gormbasics1.creditcard (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    number text,
    user_id bigint
);


ALTER TABLE gormbasics1.creditcard OWNER TO vatsal;

--
-- Name: creditcard_id_seq; Type: SEQUENCE; Schema: gormbasics1; Owner: vatsal
--

CREATE SEQUENCE gormbasics1.creditcard_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE gormbasics1.creditcard_id_seq OWNER TO vatsal;

--
-- Name: creditcard_id_seq; Type: SEQUENCE OWNED BY; Schema: gormbasics1; Owner: vatsal
--

ALTER SEQUENCE gormbasics1.creditcard_id_seq OWNED BY gormbasics1.creditcard.id;


--
-- Name: note; Type: TABLE; Schema: gormbasics1; Owner: vatsal
--

CREATE TABLE gormbasics1.note (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name character varying(255),
    content text,
    user_id bigint
);


ALTER TABLE gormbasics1.note OWNER TO vatsal;

--
-- Name: note_id_seq; Type: SEQUENCE; Schema: gormbasics1; Owner: vatsal
--

CREATE SEQUENCE gormbasics1.note_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE gormbasics1.note_id_seq OWNER TO vatsal;

--
-- Name: note_id_seq; Type: SEQUENCE OWNED BY; Schema: gormbasics1; Owner: vatsal
--

ALTER SEQUENCE gormbasics1.note_id_seq OWNED BY gormbasics1.note.id;


--
-- Name: product; Type: TABLE; Schema: gormbasics1; Owner: vatsal
--

CREATE TABLE gormbasics1.product (
    id bigint NOT NULL,
    name text,
    price bigint,
    category_name text
);


ALTER TABLE gormbasics1.product OWNER TO vatsal;

--
-- Name: product_id_seq; Type: SEQUENCE; Schema: gormbasics1; Owner: vatsal
--

CREATE SEQUENCE gormbasics1.product_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE gormbasics1.product_id_seq OWNER TO vatsal;

--
-- Name: product_id_seq; Type: SEQUENCE OWNED BY; Schema: gormbasics1; Owner: vatsal
--

ALTER SEQUENCE gormbasics1.product_id_seq OWNED BY gormbasics1.product.id;


--
-- Name: users; Type: TABLE; Schema: gormbasics1; Owner: vatsal
--

CREATE TABLE gormbasics1.users (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    username character varying(64),
    password character varying(255)
);


ALTER TABLE gormbasics1.users OWNER TO vatsal;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: gormbasics1; Owner: vatsal
--

CREATE SEQUENCE gormbasics1.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE gormbasics1.users_id_seq OWNER TO vatsal;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: gormbasics1; Owner: vatsal
--

ALTER SEQUENCE gormbasics1.users_id_seq OWNED BY gormbasics1.users.id;


--
-- Name: school; Type: TABLE; Schema: gormschool; Owner: vatsal
--

CREATE TABLE gormschool.school (
    id bigint NOT NULL,
    name text
);


ALTER TABLE gormschool.school OWNER TO vatsal;

--
-- Name: school_id_seq; Type: SEQUENCE; Schema: gormschool; Owner: vatsal
--

CREATE SEQUENCE gormschool.school_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE gormschool.school_id_seq OWNER TO vatsal;

--
-- Name: school_id_seq; Type: SEQUENCE OWNED BY; Schema: gormschool; Owner: vatsal
--

ALTER SEQUENCE gormschool.school_id_seq OWNED BY gormschool.school.id;


--
-- Name: student; Type: TABLE; Schema: gormschool; Owner: vatsal
--

CREATE TABLE gormschool.student (
    id bigint NOT NULL,
    name text,
    school_id bigint
);


ALTER TABLE gormschool.student OWNER TO vatsal;

--
-- Name: student_id_seq; Type: SEQUENCE; Schema: gormschool; Owner: vatsal
--

CREATE SEQUENCE gormschool.student_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE gormschool.student_id_seq OWNER TO vatsal;

--
-- Name: student_id_seq; Type: SEQUENCE OWNED BY; Schema: gormschool; Owner: vatsal
--

ALTER SEQUENCE gormschool.student_id_seq OWNED BY gormschool.student.id;


--
-- Name: parent; Type: TABLE; Schema: gormschoolproject; Owner: vatsal
--

CREATE TABLE gormschoolproject.parent (
    id bigint NOT NULL,
    name text,
    relation text
);


ALTER TABLE gormschoolproject.parent OWNER TO vatsal;

--
-- Name: parent_id_seq; Type: SEQUENCE; Schema: gormschoolproject; Owner: vatsal
--

CREATE SEQUENCE gormschoolproject.parent_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE gormschoolproject.parent_id_seq OWNER TO vatsal;

--
-- Name: parent_id_seq; Type: SEQUENCE OWNED BY; Schema: gormschoolproject; Owner: vatsal
--

ALTER SEQUENCE gormschoolproject.parent_id_seq OWNED BY gormschoolproject.parent.id;


--
-- Name: sport; Type: TABLE; Schema: gormschoolproject; Owner: vatsal
--

CREATE TABLE gormschoolproject.sport (
    id bigint NOT NULL,
    name text
);


ALTER TABLE gormschoolproject.sport OWNER TO vatsal;

--
-- Name: sport_id_seq; Type: SEQUENCE; Schema: gormschoolproject; Owner: vatsal
--

CREATE SEQUENCE gormschoolproject.sport_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE gormschoolproject.sport_id_seq OWNER TO vatsal;

--
-- Name: sport_id_seq; Type: SEQUENCE OWNED BY; Schema: gormschoolproject; Owner: vatsal
--

ALTER SEQUENCE gormschoolproject.sport_id_seq OWNED BY gormschoolproject.sport.id;


--
-- Name: student; Type: TABLE; Schema: gormschoolproject; Owner: vatsal
--

CREATE TABLE gormschoolproject.student (
    id bigint NOT NULL,
    name text,
    address text,
    city text,
    pincode text,
    birth_date text,
    sport_id bigint
);


ALTER TABLE gormschoolproject.student OWNER TO vatsal;

--
-- Name: student_id_seq; Type: SEQUENCE; Schema: gormschoolproject; Owner: vatsal
--

CREATE SEQUENCE gormschoolproject.student_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE gormschoolproject.student_id_seq OWNER TO vatsal;

--
-- Name: student_id_seq; Type: SEQUENCE OWNED BY; Schema: gormschoolproject; Owner: vatsal
--

ALTER SEQUENCE gormschoolproject.student_id_seq OWNED BY gormschoolproject.student.id;


--
-- Name: student_parent; Type: TABLE; Schema: gormschoolproject; Owner: vatsal
--

CREATE TABLE gormschoolproject.student_parent (
    parent_id bigint NOT NULL,
    student_id bigint NOT NULL
);


ALTER TABLE gormschoolproject.student_parent OWNER TO vatsal;

--
-- Name: user; Type: TABLE; Schema: httpnet; Owner: vatsal
--

CREATE TABLE httpnet."user" (
    id integer NOT NULL,
    firstname character varying(80) NOT NULL,
    lastname character varying(80),
    email character varying(80) NOT NULL,
    phone character varying(20),
    dateofbirth date
);


ALTER TABLE httpnet."user" OWNER TO vatsal;

--
-- Name: user_id_seq; Type: SEQUENCE; Schema: httpnet; Owner: vatsal
--

CREATE SEQUENCE httpnet.user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE httpnet.user_id_seq OWNER TO vatsal;

--
-- Name: user_id_seq; Type: SEQUENCE OWNED BY; Schema: httpnet; Owner: vatsal
--

ALTER SEQUENCE httpnet.user_id_seq OWNED BY httpnet."user".id;


--
-- Name: product_groups; Type: TABLE; Schema: products; Owner: vatsal
--

CREATE TABLE products.product_groups (
    group_id integer NOT NULL,
    group_name character varying(255) NOT NULL
);


ALTER TABLE products.product_groups OWNER TO vatsal;

--
-- Name: product_groups_group_id_seq; Type: SEQUENCE; Schema: products; Owner: vatsal
--

CREATE SEQUENCE products.product_groups_group_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE products.product_groups_group_id_seq OWNER TO vatsal;

--
-- Name: product_groups_group_id_seq; Type: SEQUENCE OWNED BY; Schema: products; Owner: vatsal
--

ALTER SEQUENCE products.product_groups_group_id_seq OWNED BY products.product_groups.group_id;


--
-- Name: products; Type: TABLE; Schema: products; Owner: vatsal
--

CREATE TABLE products.products (
    product_id integer NOT NULL,
    product_name character varying(255) NOT NULL,
    price numeric(11,2),
    group_id integer NOT NULL
);


ALTER TABLE products.products OWNER TO vatsal;

--
-- Name: products_product_id_seq; Type: SEQUENCE; Schema: products; Owner: vatsal
--

CREATE SEQUENCE products.products_product_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE products.products_product_id_seq OWNER TO vatsal;

--
-- Name: products_product_id_seq; Type: SEQUENCE OWNED BY; Schema: products; Owner: vatsal
--

ALTER SEQUENCE products.products_product_id_seq OWNED BY products.products.product_id;


--
-- Name: books; Type: TABLE; Schema: public; Owner: vatsal
--

CREATE TABLE public.books (
    id integer,
    name character varying(20),
    quantity integer
);


ALTER TABLE public.books OWNER TO vatsal;

--
-- Name: busstatus; Type: TABLE; Schema: transport; Owner: vatsal
--

CREATE TABLE transport.busstatus (
    bus_id integer NOT NULL,
    lat double precision,
    long double precision,
    last_updated time without time zone DEFAULT CURRENT_TIME NOT NULL,
    traffic integer,
    status integer DEFAULT 0 NOT NULL,
    last_station_order integer DEFAULT 0 NOT NULL
);


ALTER TABLE transport.busstatus OWNER TO vatsal;

--
-- Name: schedule; Type: TABLE; Schema: transport; Owner: vatsal
--

CREATE TABLE transport.schedule (
    id integer NOT NULL,
    bus_id integer NOT NULL,
    route_id integer NOT NULL,
    departure_time time without time zone NOT NULL
);


ALTER TABLE transport.schedule OWNER TO vatsal;

--
-- Name: businfotogether; Type: VIEW; Schema: public; Owner: vatsal
--

CREATE VIEW public.businfotogether AS
 SELECT schedule.bus_id,
    schedule.route_id,
    busstatus.last_station_order,
    busstatus.lat,
    busstatus.long,
    busstatus.last_updated,
    busstatus.traffic,
    busstatus.status
   FROM (transport.schedule
     JOIN transport.busstatus ON ((schedule.bus_id = busstatus.bus_id)));


ALTER TABLE public.businfotogether OWNER TO vatsal;

--
-- Name: route; Type: TABLE; Schema: transport; Owner: vatsal
--

CREATE TABLE transport.route (
    id integer NOT NULL,
    name character varying(10) NOT NULL,
    status integer DEFAULT 1 NOT NULL,
    source integer NOT NULL,
    destination integer NOT NULL
);


ALTER TABLE transport.route OWNER TO vatsal;

--
-- Name: routestations; Type: TABLE; Schema: transport; Owner: vatsal
--

CREATE TABLE transport.routestations (
    route_id integer NOT NULL,
    station_id integer NOT NULL,
    station_order integer NOT NULL
);


ALTER TABLE transport.routestations OWNER TO vatsal;

--
-- Name: station; Type: TABLE; Schema: transport; Owner: vatsal
--

CREATE TABLE transport.station (
    id integer NOT NULL,
    name character varying(50) NOT NULL,
    lat character varying(10) NOT NULL,
    long character varying(10) NOT NULL
);


ALTER TABLE transport.station OWNER TO vatsal;

--
-- Name: bustransport2; Type: VIEW; Schema: public; Owner: vatsal
--

CREATE VIEW public.bustransport2 AS
 SELECT v2.bus_id,
    v2.route_id,
    v2.station_id,
    v2.name AS route_name,
    v2.departure_time,
    v2.source,
    v2.destination,
    v2.status,
    v2.station_order,
    l.lat,
    l.long,
    l.name AS station_name
   FROM (( SELECT s.bus_id,
            v1.route_id,
            v1.name,
            s.departure_time,
            v1.source,
            v1.destination,
            v1.status,
            v1.station_order,
            v1.station_id
           FROM (( SELECT r.id AS route_id,
                    r.name,
                    rs.station_id,
                    r.source,
                    r.destination,
                    r.status,
                    rs.station_order
                   FROM (transport.routestations rs
                     JOIN transport.route r ON ((rs.route_id = r.id)))) v1
             JOIN transport.schedule s ON ((s.route_id = v1.route_id)))) v2
     JOIN transport.station l ON ((v2.station_id = l.id)));


ALTER TABLE public.bustransport2 OWNER TO vatsal;

--
-- Name: bustransportsystem; Type: VIEW; Schema: public; Owner: vatsal
--

CREATE VIEW public.bustransportsystem AS
 SELECT v2.bus_id,
    v2.route_id,
    v2.station_id,
    v2.name AS route_name,
    v2.departure_time,
    v2.source,
    v2.destination,
    v2.status,
    v2.station_order,
    l.lat,
    l.long,
    l.name AS station_name
   FROM (( SELECT s.bus_id,
            v1.route_id,
            v1.name,
            s.departure_time,
            v1.source,
            v1.destination,
            v1.status,
            v1.station_order,
            v1.station_id
           FROM (( SELECT r.id AS route_id,
                    r.name,
                    rs.station_id,
                    r.source,
                    r.destination,
                    r.status,
                    rs.station_order
                   FROM (transport.routestations rs
                     JOIN transport.route r ON ((rs.route_id = r.id)))) v1
             JOIN transport.schedule s ON ((s.route_id = v1.route_id)))) v2
     JOIN transport.station l ON ((v2.station_id = l.id)));


ALTER TABLE public.bustransportsystem OWNER TO vatsal;

--
-- Name: categories; Type: TABLE; Schema: public; Owner: vatsal
--

CREATE TABLE public.categories (
    name text NOT NULL,
    remark text
);


ALTER TABLE public.categories OWNER TO vatsal;

--
-- Name: companies; Type: TABLE; Schema: public; Owner: vatsal
--

CREATE TABLE public.companies (
    company_id bigint,
    code text,
    name text
);


ALTER TABLE public.companies OWNER TO vatsal;

--
-- Name: customers; Type: TABLE; Schema: public; Owner: vatsal
--

CREATE TABLE public.customers (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    first_name text,
    last_name text,
    email text,
    dateofbirth text,
    mobilenumber text
);


ALTER TABLE public.customers OWNER TO vatsal;

--
-- Name: customers_id_seq; Type: SEQUENCE; Schema: public; Owner: vatsal
--

CREATE SEQUENCE public.customers_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.customers_id_seq OWNER TO vatsal;

--
-- Name: customers_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: vatsal
--

ALTER SEQUENCE public.customers_id_seq OWNED BY public.customers.id;


--
-- Name: emp; Type: TABLE; Schema: public; Owner: vatsal
--

CREATE TABLE public.emp (
    id integer NOT NULL,
    name character varying(80),
    age integer
);


ALTER TABLE public.emp OWNER TO vatsal;

--
-- Name: emp_id_seq; Type: SEQUENCE; Schema: public; Owner: vatsal
--

CREATE SEQUENCE public.emp_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.emp_id_seq OWNER TO vatsal;

--
-- Name: emp_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: vatsal
--

ALTER SEQUENCE public.emp_id_seq OWNED BY public.emp.id;


--
-- Name: myroute; Type: VIEW; Schema: public; Owner: vatsal
--

CREATE VIEW public.myroute AS
 SELECT routestations.route_id
   FROM transport.routestations
  WHERE (routestations.station_id = 3);


ALTER TABLE public.myroute OWNER TO vatsal;

--
-- Name: product_groups; Type: TABLE; Schema: public; Owner: vatsal
--

CREATE TABLE public.product_groups (
    group_id integer NOT NULL,
    group_name character varying(255) NOT NULL
);


ALTER TABLE public.product_groups OWNER TO vatsal;

--
-- Name: product_groups_group_id_seq; Type: SEQUENCE; Schema: public; Owner: vatsal
--

CREATE SEQUENCE public.product_groups_group_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.product_groups_group_id_seq OWNER TO vatsal;

--
-- Name: product_groups_group_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: vatsal
--

ALTER SEQUENCE public.product_groups_group_id_seq OWNED BY public.product_groups.group_id;


--
-- Name: products; Type: TABLE; Schema: public; Owner: vatsal
--

CREATE TABLE public.products (
    product_id integer NOT NULL,
    product_name character varying(255) NOT NULL,
    price bigint,
    group_id integer NOT NULL,
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    code text
);


ALTER TABLE public.products OWNER TO vatsal;

--
-- Name: products_id_seq; Type: SEQUENCE; Schema: public; Owner: vatsal
--

CREATE SEQUENCE public.products_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.products_id_seq OWNER TO vatsal;

--
-- Name: products_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: vatsal
--

ALTER SEQUENCE public.products_id_seq OWNED BY public.products.id;


--
-- Name: products_product_id_seq; Type: SEQUENCE; Schema: public; Owner: vatsal
--

CREATE SEQUENCE public.products_product_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.products_product_id_seq OWNER TO vatsal;

--
-- Name: products_product_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: vatsal
--

ALTER SEQUENCE public.products_product_id_seq OWNED BY public.products.product_id;


--
-- Name: report; Type: TABLE; Schema: public; Owner: vatsal
--

CREATE TABLE public.report (
    student_id integer,
    score integer
);


ALTER TABLE public.report OWNER TO vatsal;

--
-- Name: routeshavearrivingbuses; Type: VIEW; Schema: public; Owner: vatsal
--

CREATE VIEW public.routeshavearrivingbuses AS
 SELECT schedule.route_id
   FROM transport.schedule
  WHERE ((schedule.route_id IN ( SELECT myroute.route_id
           FROM public.myroute)) AND (EXTRACT(hour FROM CURRENT_TIME) > EXTRACT(hour FROM schedule.departure_time)));


ALTER TABLE public.routeshavearrivingbuses OWNER TO vatsal;

--
-- Name: student_parent; Type: TABLE; Schema: public; Owner: vatsal
--

CREATE TABLE public.student_parent (
    parent_id bigint NOT NULL,
    student_id bigint NOT NULL
);


ALTER TABLE public.student_parent OWNER TO vatsal;

--
-- Name: studs; Type: TABLE; Schema: public; Owner: vatsal
--

CREATE TABLE public.studs (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    code text,
    price bigint
);


ALTER TABLE public.studs OWNER TO vatsal;

--
-- Name: studs_id_seq; Type: SEQUENCE; Schema: public; Owner: vatsal
--

CREATE SEQUENCE public.studs_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.studs_id_seq OWNER TO vatsal;

--
-- Name: studs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: vatsal
--

ALTER SEQUENCE public.studs_id_seq OWNED BY public.studs.id;


--
-- Name: v1; Type: VIEW; Schema: public; Owner: vatsal
--

CREATE VIEW public.v1 AS
 SELECT o.route_id,
    o.station_id,
    o.station_order
   FROM (transport.routestations o
     JOIN ( SELECT r.route_id,
            r.station_order AS myorder
           FROM transport.routestations r
          WHERE (r.station_id = 88)) s ON ((o.route_id = s.route_id)))
  WHERE (o.station_order > s.myorder);


ALTER TABLE public.v1 OWNER TO vatsal;

--
-- Name: v2; Type: VIEW; Schema: public; Owner: vatsal
--

CREATE VIEW public.v2 AS
 SELECT o.route_id,
    o.station_id,
    o.station_order
   FROM (transport.routestations o
     JOIN ( SELECT r.route_id,
            r.station_order AS myorder
           FROM transport.routestations r
          WHERE (r.station_id = 32)) s ON ((o.route_id = s.route_id)))
  WHERE (o.station_order > s.myorder);


ALTER TABLE public.v2 OWNER TO vatsal;

--
-- Name: x; Type: VIEW; Schema: public; Owner: vatsal
--

CREATE VIEW public.x AS
 SELECT route.name,
    route.source,
    route.destination
   FROM (transport.routestations
     JOIN transport.route ON ((routestations.route_id = route.id)))
  WHERE (routestations.station_id = 1);


ALTER TABLE public.x OWNER TO vatsal;

--
-- Name: class; Type: TABLE; Schema: school; Owner: vatsal
--

CREATE TABLE school.class (
    id integer NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    name character varying(255)
);


ALTER TABLE school.class OWNER TO vatsal;

--
-- Name: class_id_seq; Type: SEQUENCE; Schema: school; Owner: vatsal
--

CREATE SEQUENCE school.class_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE school.class_id_seq OWNER TO vatsal;

--
-- Name: class_id_seq; Type: SEQUENCE OWNED BY; Schema: school; Owner: vatsal
--

ALTER SEQUENCE school.class_id_seq OWNED BY school.class.id;


--
-- Name: classsubjectrelation; Type: TABLE; Schema: school; Owner: vatsal
--

CREATE TABLE school.classsubjectrelation (
    class_id integer NOT NULL,
    subject_id integer NOT NULL,
    teacher_id integer
);


ALTER TABLE school.classsubjectrelation OWNER TO vatsal;

--
-- Name: parent; Type: TABLE; Schema: school; Owner: vatsal
--

CREATE TABLE school.parent (
    id integer NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    name character varying(255),
    relation character varying(255),
    email character varying(50),
    phone character varying(12)
);


ALTER TABLE school.parent OWNER TO vatsal;

--
-- Name: parent_id_seq; Type: SEQUENCE; Schema: school; Owner: vatsal
--

CREATE SEQUENCE school.parent_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE school.parent_id_seq OWNER TO vatsal;

--
-- Name: parent_id_seq; Type: SEQUENCE OWNED BY; Schema: school; Owner: vatsal
--

ALTER SEQUENCE school.parent_id_seq OWNED BY school.parent.id;


--
-- Name: registration; Type: TABLE; Schema: school; Owner: vatsal
--

CREATE TABLE school.registration (
    student_id integer NOT NULL,
    class_id integer NOT NULL,
    year integer NOT NULL
);


ALTER TABLE school.registration OWNER TO vatsal;

--
-- Name: student_id_seq; Type: SEQUENCE; Schema: school; Owner: vatsal
--

CREATE SEQUENCE school.student_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE school.student_id_seq OWNER TO vatsal;

--
-- Name: student_id_seq; Type: SEQUENCE OWNED BY; Schema: school; Owner: vatsal
--

ALTER SEQUENCE school.student_id_seq OWNED BY school.student.id;


--
-- Name: studentparentrelation; Type: TABLE; Schema: school; Owner: vatsal
--

CREATE TABLE school.studentparentrelation (
    student_id integer NOT NULL,
    parent_id integer NOT NULL
);


ALTER TABLE school.studentparentrelation OWNER TO vatsal;

--
-- Name: subject; Type: TABLE; Schema: school; Owner: vatsal
--

CREATE TABLE school.subject (
    id integer NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    name character varying(255)
);


ALTER TABLE school.subject OWNER TO vatsal;

--
-- Name: subject_id_seq; Type: SEQUENCE; Schema: school; Owner: vatsal
--

CREATE SEQUENCE school.subject_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE school.subject_id_seq OWNER TO vatsal;

--
-- Name: subject_id_seq; Type: SEQUENCE OWNED BY; Schema: school; Owner: vatsal
--

ALTER SEQUENCE school.subject_id_seq OWNED BY school.subject.id;


--
-- Name: teacher; Type: TABLE; Schema: school; Owner: vatsal
--

CREATE TABLE school.teacher (
    id integer NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    name character varying(255)
);


ALTER TABLE school.teacher OWNER TO vatsal;

--
-- Name: teacher_id_seq; Type: SEQUENCE; Schema: school; Owner: vatsal
--

CREATE SEQUENCE school.teacher_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE school.teacher_id_seq OWNER TO vatsal;

--
-- Name: teacher_id_seq; Type: SEQUENCE OWNED BY; Schema: school; Owner: vatsal
--

ALTER SEQUENCE school.teacher_id_seq OWNED BY school.teacher.id;


--
-- Name: emp; Type: TABLE; Schema: sql; Owner: vatsal
--

CREATE TABLE sql.emp (
    id integer NOT NULL,
    name character varying(80),
    age integer
);


ALTER TABLE sql.emp OWNER TO vatsal;

--
-- Name: emp_id_seq; Type: SEQUENCE; Schema: sql; Owner: vatsal
--

CREATE SEQUENCE sql.emp_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sql.emp_id_seq OWNER TO vatsal;

--
-- Name: emp_id_seq; Type: SEQUENCE OWNED BY; Schema: sql; Owner: vatsal
--

ALTER SEQUENCE sql.emp_id_seq OWNED BY sql.emp.id;


--
-- Name: books; Type: TABLE; Schema: task6muxgorm; Owner: vatsal
--

CREATE TABLE task6muxgorm.books (
    id bigint NOT NULL,
    title text,
    author text,
    isbn text,
    publisher text,
    year bigint,
    genre text,
    quantity bigint
);


ALTER TABLE task6muxgorm.books OWNER TO vatsal;

--
-- Name: books_id_seq; Type: SEQUENCE; Schema: task6muxgorm; Owner: vatsal
--

CREATE SEQUENCE task6muxgorm.books_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE task6muxgorm.books_id_seq OWNER TO vatsal;

--
-- Name: books_id_seq; Type: SEQUENCE OWNED BY; Schema: task6muxgorm; Owner: vatsal
--

ALTER SEQUENCE task6muxgorm.books_id_seq OWNED BY task6muxgorm.books.id;


--
-- Name: users; Type: TABLE; Schema: task6muxgorm; Owner: vatsal
--

CREATE TABLE task6muxgorm.users (
    id bigint NOT NULL,
    name text,
    password text,
    role text
);


ALTER TABLE task6muxgorm.users OWNER TO vatsal;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: task6muxgorm; Owner: vatsal
--

CREATE SEQUENCE task6muxgorm.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE task6muxgorm.users_id_seq OWNER TO vatsal;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: task6muxgorm; Owner: vatsal
--

ALTER SEQUENCE task6muxgorm.users_id_seq OWNED BY task6muxgorm.users.id;


--
-- Name: usersbooks; Type: TABLE; Schema: task6muxgorm; Owner: vatsal
--

CREATE TABLE task6muxgorm.usersbooks (
    user_id bigint NOT NULL,
    book_id bigint NOT NULL
);


ALTER TABLE task6muxgorm.usersbooks OWNER TO vatsal;

--
-- Name: todos; Type: TABLE; Schema: todo_cli; Owner: vatsal
--

CREATE TABLE todo_cli.todos (
    id integer NOT NULL,
    title character varying(255),
    completed boolean,
    createdat date
);


ALTER TABLE todo_cli.todos OWNER TO vatsal;

--
-- Name: todos_id_seq; Type: SEQUENCE; Schema: todo_cli; Owner: vatsal
--

CREATE SEQUENCE todo_cli.todos_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE todo_cli.todos_id_seq OWNER TO vatsal;

--
-- Name: todos_id_seq; Type: SEQUENCE OWNED BY; Schema: todo_cli; Owner: vatsal
--

ALTER SEQUENCE todo_cli.todos_id_seq OWNED BY todo_cli.todos.id;


--
-- Name: bus; Type: TABLE; Schema: transport; Owner: vatsal
--

CREATE TABLE transport.bus (
    id integer NOT NULL,
    registration_number character varying(10) NOT NULL,
    model character varying(15) NOT NULL,
    capacity integer NOT NULL
);


ALTER TABLE transport.bus OWNER TO vatsal;

--
-- Name: bus_id_seq; Type: SEQUENCE; Schema: transport; Owner: vatsal
--

CREATE SEQUENCE transport.bus_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE transport.bus_id_seq OWNER TO vatsal;

--
-- Name: bus_id_seq; Type: SEQUENCE OWNED BY; Schema: transport; Owner: vatsal
--

ALTER SEQUENCE transport.bus_id_seq OWNED BY transport.bus.id;


--
-- Name: driver; Type: TABLE; Schema: transport; Owner: vatsal
--

CREATE TABLE transport.driver (
    id integer NOT NULL,
    name character varying(30) NOT NULL,
    phone character varying(14) NOT NULL,
    gender integer DEFAULT 0 NOT NULL,
    dob date NOT NULL
);


ALTER TABLE transport.driver OWNER TO vatsal;

--
-- Name: driver_id_seq; Type: SEQUENCE; Schema: transport; Owner: vatsal
--

CREATE SEQUENCE transport.driver_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE transport.driver_id_seq OWNER TO vatsal;

--
-- Name: driver_id_seq; Type: SEQUENCE OWNED BY; Schema: transport; Owner: vatsal
--

ALTER SEQUENCE transport.driver_id_seq OWNED BY transport.driver.id;


--
-- Name: route_id_seq; Type: SEQUENCE; Schema: transport; Owner: vatsal
--

CREATE SEQUENCE transport.route_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE transport.route_id_seq OWNER TO vatsal;

--
-- Name: route_id_seq; Type: SEQUENCE OWNED BY; Schema: transport; Owner: vatsal
--

ALTER SEQUENCE transport.route_id_seq OWNED BY transport.route.id;


--
-- Name: schedule_id_seq; Type: SEQUENCE; Schema: transport; Owner: vatsal
--

CREATE SEQUENCE transport.schedule_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE transport.schedule_id_seq OWNER TO vatsal;

--
-- Name: schedule_id_seq; Type: SEQUENCE OWNED BY; Schema: transport; Owner: vatsal
--

ALTER SEQUENCE transport.schedule_id_seq OWNED BY transport.schedule.id;


--
-- Name: station_id_seq; Type: SEQUENCE; Schema: transport; Owner: vatsal
--

CREATE SEQUENCE transport.station_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE transport.station_id_seq OWNER TO vatsal;

--
-- Name: station_id_seq; Type: SEQUENCE OWNED BY; Schema: transport; Owner: vatsal
--

ALTER SEQUENCE transport.station_id_seq OWNED BY transport.station.id;


--
-- Name: schools id; Type: DEFAULT; Schema: arrays; Owner: vatsal
--

ALTER TABLE ONLY arrays.schools ALTER COLUMN id SET DEFAULT nextval('arrays.schools_id_seq'::regclass);


--
-- Name: students id; Type: DEFAULT; Schema: gormbasics; Owner: vatsal
--

ALTER TABLE ONLY gormbasics.students ALTER COLUMN id SET DEFAULT nextval('gormbasics.students_id_seq'::regclass);


--
-- Name: creditcard id; Type: DEFAULT; Schema: gormbasics1; Owner: vatsal
--

ALTER TABLE ONLY gormbasics1.creditcard ALTER COLUMN id SET DEFAULT nextval('gormbasics1.creditcard_id_seq'::regclass);


--
-- Name: note id; Type: DEFAULT; Schema: gormbasics1; Owner: vatsal
--

ALTER TABLE ONLY gormbasics1.note ALTER COLUMN id SET DEFAULT nextval('gormbasics1.note_id_seq'::regclass);


--
-- Name: product id; Type: DEFAULT; Schema: gormbasics1; Owner: vatsal
--

ALTER TABLE ONLY gormbasics1.product ALTER COLUMN id SET DEFAULT nextval('gormbasics1.product_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: gormbasics1; Owner: vatsal
--

ALTER TABLE ONLY gormbasics1.users ALTER COLUMN id SET DEFAULT nextval('gormbasics1.users_id_seq'::regclass);


--
-- Name: school id; Type: DEFAULT; Schema: gormschool; Owner: vatsal
--

ALTER TABLE ONLY gormschool.school ALTER COLUMN id SET DEFAULT nextval('gormschool.school_id_seq'::regclass);


--
-- Name: student id; Type: DEFAULT; Schema: gormschool; Owner: vatsal
--

ALTER TABLE ONLY gormschool.student ALTER COLUMN id SET DEFAULT nextval('gormschool.student_id_seq'::regclass);


--
-- Name: parent id; Type: DEFAULT; Schema: gormschoolproject; Owner: vatsal
--

ALTER TABLE ONLY gormschoolproject.parent ALTER COLUMN id SET DEFAULT nextval('gormschoolproject.parent_id_seq'::regclass);


--
-- Name: sport id; Type: DEFAULT; Schema: gormschoolproject; Owner: vatsal
--

ALTER TABLE ONLY gormschoolproject.sport ALTER COLUMN id SET DEFAULT nextval('gormschoolproject.sport_id_seq'::regclass);


--
-- Name: student id; Type: DEFAULT; Schema: gormschoolproject; Owner: vatsal
--

ALTER TABLE ONLY gormschoolproject.student ALTER COLUMN id SET DEFAULT nextval('gormschoolproject.student_id_seq'::regclass);


--
-- Name: user id; Type: DEFAULT; Schema: httpnet; Owner: vatsal
--

ALTER TABLE ONLY httpnet."user" ALTER COLUMN id SET DEFAULT nextval('httpnet.user_id_seq'::regclass);


--
-- Name: product_groups group_id; Type: DEFAULT; Schema: products; Owner: vatsal
--

ALTER TABLE ONLY products.product_groups ALTER COLUMN group_id SET DEFAULT nextval('products.product_groups_group_id_seq'::regclass);


--
-- Name: products product_id; Type: DEFAULT; Schema: products; Owner: vatsal
--

ALTER TABLE ONLY products.products ALTER COLUMN product_id SET DEFAULT nextval('products.products_product_id_seq'::regclass);


--
-- Name: customers id; Type: DEFAULT; Schema: public; Owner: vatsal
--

ALTER TABLE ONLY public.customers ALTER COLUMN id SET DEFAULT nextval('public.customers_id_seq'::regclass);


--
-- Name: emp id; Type: DEFAULT; Schema: public; Owner: vatsal
--

ALTER TABLE ONLY public.emp ALTER COLUMN id SET DEFAULT nextval('public.emp_id_seq'::regclass);


--
-- Name: product_groups group_id; Type: DEFAULT; Schema: public; Owner: vatsal
--

ALTER TABLE ONLY public.product_groups ALTER COLUMN group_id SET DEFAULT nextval('public.product_groups_group_id_seq'::regclass);


--
-- Name: products product_id; Type: DEFAULT; Schema: public; Owner: vatsal
--

ALTER TABLE ONLY public.products ALTER COLUMN product_id SET DEFAULT nextval('public.products_product_id_seq'::regclass);


--
-- Name: products id; Type: DEFAULT; Schema: public; Owner: vatsal
--

ALTER TABLE ONLY public.products ALTER COLUMN id SET DEFAULT nextval('public.products_id_seq'::regclass);


--
-- Name: studs id; Type: DEFAULT; Schema: public; Owner: vatsal
--

ALTER TABLE ONLY public.studs ALTER COLUMN id SET DEFAULT nextval('public.studs_id_seq'::regclass);


--
-- Name: class id; Type: DEFAULT; Schema: school; Owner: vatsal
--

ALTER TABLE ONLY school.class ALTER COLUMN id SET DEFAULT nextval('school.class_id_seq'::regclass);


--
-- Name: parent id; Type: DEFAULT; Schema: school; Owner: vatsal
--

ALTER TABLE ONLY school.parent ALTER COLUMN id SET DEFAULT nextval('school.parent_id_seq'::regclass);


--
-- Name: subject id; Type: DEFAULT; Schema: school; Owner: vatsal
--

ALTER TABLE ONLY school.subject ALTER COLUMN id SET DEFAULT nextval('school.subject_id_seq'::regclass);


--
-- Name: teacher id; Type: DEFAULT; Schema: school; Owner: vatsal
--

ALTER TABLE ONLY school.teacher ALTER COLUMN id SET DEFAULT nextval('school.teacher_id_seq'::regclass);


--
-- Name: emp id; Type: DEFAULT; Schema: sql; Owner: vatsal
--

ALTER TABLE ONLY sql.emp ALTER COLUMN id SET DEFAULT nextval('sql.emp_id_seq'::regclass);


--
-- Name: books id; Type: DEFAULT; Schema: task6muxgorm; Owner: vatsal
--

ALTER TABLE ONLY task6muxgorm.books ALTER COLUMN id SET DEFAULT nextval('task6muxgorm.books_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: task6muxgorm; Owner: vatsal
--

ALTER TABLE ONLY task6muxgorm.users ALTER COLUMN id SET DEFAULT nextval('task6muxgorm.users_id_seq'::regclass);


--
-- Name: todos id; Type: DEFAULT; Schema: todo_cli; Owner: vatsal
--

ALTER TABLE ONLY todo_cli.todos ALTER COLUMN id SET DEFAULT nextval('todo_cli.todos_id_seq'::regclass);


--
-- Name: driver id; Type: DEFAULT; Schema: transport; Owner: vatsal
--

ALTER TABLE ONLY transport.driver ALTER COLUMN id SET DEFAULT nextval('transport.driver_id_seq'::regclass);


--
-- Data for Name: schools; Type: TABLE DATA; Schema: arrays; Owner: vatsal
--

COPY arrays.schools (id, names, ranks) FROM stdin;
1	{John,Alice,Bob}	{1,2,3,4,5}
2	{Mary,David}	{10,20,30}
3	{John,Alice,Bob}	{1,2,3,4,5}
4	{Mary,David}	{10,20,30}
\.


--
-- Data for Name: generate_column; Type: TABLE DATA; Schema: data_definition; Owner: vatsal
--

COPY data_definition.generate_column (data_km) FROM stdin;
45
\.


--
-- Data for Name: deestinct; Type: TABLE DATA; Schema: doc5; Owner: vatsal
--

COPY doc5.deestinct (id, name) FROM stdin;
1	good
2	good
\.


--
-- Data for Name: students; Type: TABLE DATA; Schema: gormbasics; Owner: vatsal
--

COPY gormbasics.students (id, created_at, updated_at, deleted_at, name, class, age) FROM stdin;
1	2024-02-14 13:10:46.010199+05:30	2024-02-14 13:10:46.010199+05:30	\N	Aman	12	\N
2	2024-02-14 13:10:53.321263+05:30	2024-02-14 13:10:53.321263+05:30	\N	Aman	12	\N
3	2024-02-14 13:19:08.902274+05:30	2024-02-14 13:19:08.902274+05:30	\N	Aman	12	\N
4	2024-02-14 13:19:27.925406+05:30	2024-02-14 13:19:27.925406+05:30	\N	Aman	12	\N
5	2024-02-14 13:23:24.707756+05:30	2024-02-14 13:23:24.707756+05:30	\N	Aman	12	\N
6	2024-02-14 13:23:34.689076+05:30	2024-02-14 13:23:34.689076+05:30	\N	Aman	12	\N
7	2024-02-14 13:23:57.202276+05:30	2024-02-14 13:23:57.202276+05:30	\N	Ketan	12	\N
8	2024-02-14 13:26:21.791126+05:30	2024-02-14 13:26:21.791126+05:30	\N	Ketan	12	\N
9	2024-02-14 13:26:41.849292+05:30	2024-02-14 13:26:41.849292+05:30	\N	Ketan	12	\N
10	2024-02-14 13:26:56.243004+05:30	2024-02-14 13:26:56.243004+05:30	\N	Ketan	12	\N
11	2024-02-14 13:26:56.292877+05:30	2024-02-14 13:26:56.292877+05:30	\N	Ketan	\N	\N
12	2024-02-14 13:32:36.394066+05:30	2024-02-14 13:32:36.394066+05:30	\N	Ketan	12	\N
13	2024-02-14 13:32:36.457633+05:30	2024-02-14 13:32:36.457633+05:30	\N	Ketan	\N	\N
14	2024-02-14 13:32:36.487296+05:30	2024-02-14 13:32:36.487296+05:30	\N	Aman	9	\N
15	2024-02-14 13:32:36.487296+05:30	2024-02-14 13:32:36.487296+05:30	\N	Rahul	30	\N
16	2024-02-14 13:32:36.488016+05:30	2024-02-14 13:32:36.488016+05:30	\N	Rahul Sharma	0	\N
17	2024-02-14 13:32:36.488016+05:30	2024-02-14 13:32:36.488016+05:30	\N	Tridip Chavda	30	\N
18	2024-02-14 13:32:36.488228+05:30	2024-02-14 13:32:36.488228+05:30	\N	Tridip Chavda 2	20	\N
19	2024-02-14 13:33:29.896864+05:30	2024-02-14 13:33:29.896864+05:30	\N	Ketan	12	\N
20	2024-02-14 13:33:29.938348+05:30	2024-02-14 13:33:29.938348+05:30	\N	Ketan	\N	\N
21	2024-02-14 13:33:29.98762+05:30	2024-02-14 13:33:29.98762+05:30	\N	Aman	9	\N
22	2024-02-14 13:33:29.98762+05:30	2024-02-14 13:33:29.98762+05:30	\N	Rahul	30	\N
23	2024-02-14 13:33:29.988017+05:30	2024-02-14 13:33:29.988017+05:30	\N	Rahul Sharma	0	\N
24	2024-02-14 13:33:29.988017+05:30	2024-02-14 13:33:29.988017+05:30	\N	Tridip Chavda	30	\N
25	2024-02-14 13:33:29.988189+05:30	2024-02-14 13:33:29.988189+05:30	\N	Tridip Chavda 2	20	\N
26	2024-02-14 13:37:17.288774+05:30	2024-02-14 13:37:17.288774+05:30	\N	Ketan	12	\N
27	2024-02-14 13:37:17.335469+05:30	2024-02-14 13:37:17.335469+05:30	\N	Ketan	\N	\N
28	2024-02-14 13:37:17.385287+05:30	2024-02-14 13:37:17.385287+05:30	\N	Aman	9	\N
29	2024-02-14 13:37:17.385287+05:30	2024-02-14 13:37:17.385287+05:30	\N	Rahul	30	\N
30	2024-02-14 13:37:17.386215+05:30	2024-02-14 13:37:17.386215+05:30	\N	Rahul Sharma	0	\N
31	2024-02-14 13:37:17.386215+05:30	2024-02-14 13:37:17.386215+05:30	\N	Tridip Chavda	30	\N
32	2024-02-14 13:37:17.386773+05:30	2024-02-14 13:37:17.386773+05:30	\N	Tridip Chavda 2	20	\N
33	2024-02-14 13:38:23.230445+05:30	2024-02-14 13:38:23.230445+05:30	\N	Ketan	12	\N
34	2024-02-14 13:38:23.27611+05:30	2024-02-14 13:38:23.27611+05:30	\N	Ketan	\N	\N
35	2024-02-14 13:39:06.743999+05:30	2024-02-14 13:39:06.743999+05:30	\N	Ketan	12	\N
36	2024-02-14 13:39:06.772555+05:30	2024-02-14 13:39:06.772555+05:30	\N	Ketan	\N	\N
37	2024-02-14 13:40:31.828618+05:30	2024-02-14 13:40:31.828618+05:30	\N	Ketan	12	\N
38	2024-02-14 13:40:31.876479+05:30	2024-02-14 13:40:31.876479+05:30	\N	Ketan	\N	\N
39	2024-02-14 13:40:31.924501+05:30	2024-02-14 13:40:31.924501+05:30	\N	AMAN created	20	\N
40	2024-02-14 13:50:36.343727+05:30	2024-02-14 13:50:36.343727+05:30	\N	Ketan	12	\N
41	2024-02-14 13:50:36.424827+05:30	2024-02-14 13:50:36.424827+05:30	\N	Ketan	\N	\N
42	2024-02-14 13:50:36.449539+05:30	2024-02-14 13:50:36.449539+05:30	\N	AMAN created	20	\N
43	\N	\N	\N	jinhu	12	\N
44	2024-02-14 13:59:26.393009+05:30	2024-02-14 13:59:26.393009+05:30	\N	Ketan	12	\N
45	2024-02-14 13:59:26.465929+05:30	2024-02-14 13:59:26.465929+05:30	\N	Ketan	\N	\N
46	2024-02-14 13:59:26.515292+05:30	2024-02-14 13:59:26.515292+05:30	\N	AMAN created	20	\N
47	\N	\N	\N	jinhu	12	\N
48	2024-02-14 14:00:27.009594+05:30	2024-02-14 14:00:27.009594+05:30	\N	Ketan	12	\N
49	2024-02-14 14:00:27.056232+05:30	2024-02-14 14:00:27.056232+05:30	\N	Ketan	18	\N
50	2024-02-14 14:00:27.105337+05:30	2024-02-14 14:00:27.105337+05:30	\N	AMAN created	20	\N
51	\N	\N	\N	jinhu	12	\N
52	2024-02-14 14:01:26.610525+05:30	2024-02-14 14:01:26.610525+05:30	\N	Ketan	12	\N
53	2024-02-14 14:01:26.683197+05:30	2024-02-14 14:01:26.683197+05:30	\N	Ketan	18	\N
54	2024-02-14 14:01:26.757725+05:30	2024-02-14 14:01:26.757725+05:30	\N	AMAN created	20	\N
55	\N	\N	\N	jinhu	12	\N
56	2024-02-14 14:01:39.033308+05:30	2024-02-14 14:01:39.033308+05:30	\N	Ketan	12	\N
57	2024-02-14 14:01:39.107429+05:30	2024-02-14 14:01:39.107429+05:30	\N	Ketan	18	\N
58	2024-02-14 14:01:39.239415+05:30	2024-02-14 14:01:39.239415+05:30	\N	AMAN created	20	\N
59	\N	\N	\N	jinhu	12	\N
60	2024-02-14 14:04:50.930713+05:30	2024-02-14 14:04:50.930713+05:30	\N	Ketan	12	5
61	2024-02-14 14:04:50.963639+05:30	2024-02-14 14:04:50.963639+05:30	\N	Ketan	18	5
62	2024-02-14 14:04:50.988459+05:30	2024-02-14 14:04:50.988459+05:30	\N	AMAN created	20	5
63	\N	\N	\N	jinhu	12	5
64	2024-02-14 14:04:51.037412+05:30	2024-02-14 14:04:51.037412+05:30	\N		18	10
65	2024-02-14 14:05:12.41956+05:30	2024-02-14 14:05:12.41956+05:30	\N	Ketan	12	5
66	2024-02-14 14:05:12.45402+05:30	2024-02-14 14:05:12.45402+05:30	\N	Ketan	18	5
67	2024-02-14 14:05:12.527819+05:30	2024-02-14 14:05:12.527819+05:30	\N	AMAN created	20	5
68	\N	\N	\N	jinhu	12	5
69	2024-02-14 14:05:12.626151+05:30	2024-02-14 14:05:12.626151+05:30	\N		18	10
70	2024-02-14 14:05:12.67538+05:30	2024-02-14 14:05:12.67538+05:30	\N		18	5
71	2024-02-14 14:06:32.118054+05:30	2024-02-14 14:06:32.118054+05:30	\N	Ketan	12	5
72	2024-02-14 14:06:32.15328+05:30	2024-02-14 14:06:32.15328+05:30	\N	Ketan	18	5
73	2024-02-14 14:06:32.202496+05:30	2024-02-14 14:06:32.202496+05:30	\N	AMAN created	20	5
74	\N	\N	\N	jinhu	12	5
75	2024-02-14 14:06:32.301258+05:30	2024-02-14 14:06:32.301258+05:30	\N		18	0
76	2024-02-14 14:06:32.350454+05:30	2024-02-14 14:06:32.350454+05:30	\N		18	5
\.


--
-- Data for Name: category; Type: TABLE DATA; Schema: gormbasics1; Owner: vatsal
--

COPY gormbasics1.category (name, remark) FROM stdin;
grocery	this is the best grocery ha ha
electronic	this is the best grocery ha ha
\.


--
-- Data for Name: creditcard; Type: TABLE DATA; Schema: gormbasics1; Owner: vatsal
--

COPY gormbasics1.creditcard (id, created_at, updated_at, deleted_at, number, user_id) FROM stdin;
\.


--
-- Data for Name: note; Type: TABLE DATA; Schema: gormbasics1; Owner: vatsal
--

COPY gormbasics1.note (id, created_at, updated_at, deleted_at, name, content, user_id) FROM stdin;
1	2024-02-15 15:20:33.839798+05:30	2024-02-15 15:20:33.839798+05:30	\N			1
2	2024-02-15 15:20:35.248146+05:30	2024-02-15 15:20:35.248146+05:30	\N			1
3	2024-02-15 15:20:36.538407+05:30	2024-02-15 15:20:36.538407+05:30	\N			1
\.


--
-- Data for Name: product; Type: TABLE DATA; Schema: gormbasics1; Owner: vatsal
--

COPY gormbasics1.product (id, name, price, category_name) FROM stdin;
2	head phone 3	40	grocery
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: gormbasics1; Owner: vatsal
--

COPY gormbasics1.users (id, created_at, updated_at, deleted_at, username, password) FROM stdin;
1	2024-02-15 15:19:35.197396+05:30	2024-02-15 15:19:35.197396+05:30	\N	bacancy	
\.


--
-- Data for Name: school; Type: TABLE DATA; Schema: gormschool; Owner: vatsal
--

COPY gormschool.school (id, name) FROM stdin;
1	Archana Vidya Sankul
2	Arpan vidya sankull
\.


--
-- Data for Name: student; Type: TABLE DATA; Schema: gormschool; Owner: vatsal
--

COPY gormschool.student (id, name, school_id) FROM stdin;
1	Aman	1
2	Drishti	1
3	Rahul	1
4	Ketan	2
5	Rohan	2
\.


--
-- Data for Name: parent; Type: TABLE DATA; Schema: gormschoolproject; Owner: vatsal
--

COPY gormschoolproject.parent (id, name, relation) FROM stdin;
1	parent 1	father
2	parent 2	mother
3	parent 3	mother
\.


--
-- Data for Name: sport; Type: TABLE DATA; Schema: gormschoolproject; Owner: vatsal
--

COPY gormschoolproject.sport (id, name) FROM stdin;
1	volleball
2	volleball
3	kabbadi
4	hockey
5	table - tennis
\.


--
-- Data for Name: student; Type: TABLE DATA; Schema: gormschoolproject; Owner: vatsal
--

COPY gormschoolproject.student (id, name, address, city, pincode, birth_date, sport_id) FROM stdin;
2341	aman					1
2342	aman					1
2343	aman					1
2344	aman					1
\.


--
-- Data for Name: student_parent; Type: TABLE DATA; Schema: gormschoolproject; Owner: vatsal
--

COPY gormschoolproject.student_parent (parent_id, student_id) FROM stdin;
1	2341
2	2341
3	2342
\.


--
-- Data for Name: user; Type: TABLE DATA; Schema: httpnet; Owner: vatsal
--

COPY httpnet."user" (id, firstname, lastname, email, phone, dateofbirth) FROM stdin;
1	Ketan	Rathod	ketanrtd1@gmail.com	9725488060	2003-01-07
5	asd	sd	ketan.rathod@bacancy.com	97254880601	2024-02-01
9	Ketan	Rathod	ketan.rathod2@bacancy.com	97254880602	2024-02-01
15	Ketan	Rathod	ketan.rathod@bacancy.comee	9725488060112	2024-02-02
17	Ketan	Rathod	ketan.rathod12@bacancy.com	9725488060123	2024-02-01
19	Ketan	Rathod	ketan.rathod123@bacancy.com	97254880601234	2024-02-01
20	Ketan	Rathod	ketan.rathod21@bacancy.com	19725488060	2024-02-03
21	Ketan		ketan.rathod1234@bacancy.com	129725488060	2024-02-01
22	Tridip	Chavda	tridip@gmail.com	1234567890	2003-02-10
24	Tridip	Chavda	tridip2@gmail.com	12345678902	2003-02-10
\.


--
-- Data for Name: product_groups; Type: TABLE DATA; Schema: products; Owner: vatsal
--

COPY products.product_groups (group_id, group_name) FROM stdin;
1	Smartphone
2	Laptop
3	Tablet
\.


--
-- Data for Name: products; Type: TABLE DATA; Schema: products; Owner: vatsal
--

COPY products.products (product_id, product_name, price, group_id) FROM stdin;
1	Microsoft Lumia	200.00	1
2	HTC One	400.00	1
3	Nexus	500.00	1
4	iPhone	900.00	1
5	HP Elite	1200.00	2
6	Lenovo Thinkpad	700.00	2
7	Sony VAIO	700.00	2
8	Dell Vostro	800.00	2
9	iPad	700.00	3
10	Kindle Fire	150.00	3
11	Samsung Galaxy Tab	200.00	3
\.


--
-- Data for Name: books; Type: TABLE DATA; Schema: public; Owner: vatsal
--

COPY public.books (id, name, quantity) FROM stdin;
1	ketan	4
\.


--
-- Data for Name: categories; Type: TABLE DATA; Schema: public; Owner: vatsal
--

COPY public.categories (name, remark) FROM stdin;
grocery	It is used for home purpose
\.


--
-- Data for Name: companies; Type: TABLE DATA; Schema: public; Owner: vatsal
--

COPY public.companies (company_id, code, name) FROM stdin;
\.


--
-- Data for Name: customers; Type: TABLE DATA; Schema: public; Owner: vatsal
--

COPY public.customers (id, created_at, updated_at, deleted_at, first_name, last_name, email, dateofbirth, mobilenumber) FROM stdin;
1	2024-02-13 22:35:58.487137+05:30	2024-02-13 22:35:58.487137+05:30	\N					
2	2024-02-13 22:36:09.748065+05:30	2024-02-13 22:36:09.748065+05:30	\N					
3	2024-02-13 22:36:10.750248+05:30	2024-02-13 22:36:10.750248+05:30	\N					
4	2024-02-13 22:36:25.568187+05:30	2024-02-13 22:36:25.568187+05:30	\N	Aman shukla				
5	2024-02-13 22:36:50.502567+05:30	2024-02-13 22:36:50.502567+05:30	\N	Aman shukla	shukla	aman@gmail.com		
6	2024-02-14 12:02:10.102566+05:30	2024-02-14 12:02:10.102566+05:30	\N	Aman shukla	shukla	aman@gmail.com		
\.


--
-- Data for Name: emp; Type: TABLE DATA; Schema: public; Owner: vatsal
--

COPY public.emp (id, name, age) FROM stdin;
\.


--
-- Data for Name: product_groups; Type: TABLE DATA; Schema: public; Owner: vatsal
--

COPY public.product_groups (group_id, group_name) FROM stdin;
1	Smartphone
2	Laptop
3	Tablet
\.


--
-- Data for Name: products; Type: TABLE DATA; Schema: public; Owner: vatsal
--

COPY public.products (product_id, product_name, price, group_id, id, created_at, updated_at, deleted_at, code) FROM stdin;
1	Microsoft Lumia	200	1	1	\N	\N	\N	\N
2	HTC One	400	1	2	\N	\N	\N	\N
3	Nexus	500	1	3	\N	\N	\N	\N
4	iPhone	900	1	4	\N	\N	\N	\N
5	HP Elite	1200	2	5	\N	\N	\N	\N
6	Lenovo Thinkpad	700	2	6	\N	\N	\N	\N
7	Sony VAIO	700	2	7	\N	\N	\N	\N
8	Dell Vostro	800	2	8	\N	\N	\N	\N
9	iPad	700	3	9	\N	\N	\N	\N
10	Kindle Fire	150	3	10	\N	\N	\N	\N
11	Samsung Galaxy Tab	200	3	11	\N	\N	\N	\N
\.


--
-- Data for Name: report; Type: TABLE DATA; Schema: public; Owner: vatsal
--

COPY public.report (student_id, score) FROM stdin;
\.


--
-- Data for Name: student_parent; Type: TABLE DATA; Schema: public; Owner: vatsal
--

COPY public.student_parent (parent_id, student_id) FROM stdin;
1	2341
\.


--
-- Data for Name: studs; Type: TABLE DATA; Schema: public; Owner: vatsal
--

COPY public.studs (id, created_at, updated_at, deleted_at, code, price) FROM stdin;
1	2024-02-14 11:58:02.002588+05:30	2024-02-14 11:58:02.002588+05:30	\N	D42	100
2	2024-02-14 11:58:22.381075+05:30	2024-02-14 11:58:22.381075+05:30	\N	D42	100
3	2024-02-14 11:58:24.370542+05:30	2024-02-14 11:58:24.370542+05:30	\N	D42	100
\.


--
-- Data for Name: class; Type: TABLE DATA; Schema: school; Owner: vatsal
--

COPY school.class (id, created_at, name) FROM stdin;
1	2024-02-07 15:00:45.74188	CBSE Standard 8
2	2024-02-07 15:00:51.593395	CBSE Standard 9
\.


--
-- Data for Name: classsubjectrelation; Type: TABLE DATA; Schema: school; Owner: vatsal
--

COPY school.classsubjectrelation (class_id, subject_id, teacher_id) FROM stdin;
2	7	3
2	8	3
1	6	\N
1	7	\N
\.


--
-- Data for Name: parent; Type: TABLE DATA; Schema: school; Owner: vatsal
--

COPY school.parent (id, created_at, name, relation, email, phone) FROM stdin;
2	2024-02-07 13:56:28.772181	Aman	Father	aman@gmail.com	2345678901
3	2024-02-07 13:56:34.808204	Loveme	Father	loveme@gmail.com	9725488060
4	2024-02-07 14:17:17.372442	Bhumika	Mother	bhumika@gmail.com	9099726655
6	2024-02-09 09:48:01.619389	Anjali	Mother	anjali@gmail.com	1234567890
\.


--
-- Data for Name: registration; Type: TABLE DATA; Schema: school; Owner: vatsal
--

COPY school.registration (student_id, class_id, year) FROM stdin;
1	1	2020
2	1	2020
3	1	2020
4	1	2020
5	1	2020
6	1	2020
7	2	2020
8	2	2020
9	2	2020
10	2	2020
\.


--
-- Data for Name: student; Type: TABLE DATA; Schema: school; Owner: vatsal
--

COPY school.student (id, created_at, name, address, city, pincode, birth_date) FROM stdin;
1	2023-12-06 07:51:27	Lilyan	\N	ahmedabad	\N	2003-01-07
2	2023-12-19 19:33:44	Griffin	\N	ahmedabad	\N	2001-08-14
4	2023-07-21 03:21:45	Dre	\N	surat	\N	2001-08-14
5	2023-05-26 10:18:38	Rebeca	\N	surat	\N	2001-08-14
6	2023-10-06 19:09:32	Yolande	\N	ahmedabad	\N	2001-08-14
7	2023-10-19 22:57:57	Angeli	\N	surat	\N	2001-08-14
8	2023-03-06 03:23:11	Lenci	\N	ahmedabad	\N	2001-08-14
9	2023-08-01 08:22:17	Vivia	\N	surat	\N	2001-08-14
10	2023-07-18 07:32:55	Ignace	\N	surat	\N	2001-08-14
11	2023-10-04 12:48:26	Geordie	\N	valsad	\N	2001-02-14
12	2023-07-21 15:53:56	Clea	\N	bharuch	\N	2001-02-14
13	2023-04-10 09:01:57	Adara	\N	vadodara	\N	2001-02-14
14	2023-11-19 06:08:58	Austen	\N	bharuch	\N	2002-10-01
15	2023-11-25 22:36:25	Rorie	\N	vadodara	\N	2002-10-01
16	2024-01-25 08:16:52	Janette	\N	surendranagar	\N	2002-10-01
17	2024-01-23 00:52:56	Maria	\N	bhavnagar	\N	2002-10-01
18	2023-06-16 07:00:53	Scott	\N	bhavnagar	\N	2002-10-01
19	2023-11-25 14:01:34	Constancy	\N	amreli	\N	2002-10-01
20	2023-10-27 02:07:22	Bobinette	\N	amreli	\N	2002-10-01
3	2023-06-04 22:16:58	Berri	\N	ahmedabad	\N	2000-08-14
21	2024-02-08 12:06:23.88332	Ketan	C-33 sarthak row house valthan	surat	394325	2003-07-01
22	2024-02-08 12:33:56.583236	Aman	D-34 sarthak row house valthan	kanpur	394325	2001-12-05
27	2024-02-09 10:44:43.093019	AMAN	\N	\N	\N	2001-10-12
28	2024-02-09 10:45:22.183652	AMAN	\N	\N	\N	\N
40	2024-02-09 11:21:44.525619	afgafdga	\N	\N	\N	\N
41	2024-02-09 11:22:22.189255	afgafdga	\N	\N	\N	\N
42	2024-02-09 11:23:19.667883	afgafdga	\N	\N	\N	\N
44	2024-02-09 11:25:13.086452	afgafdga	\N	\N	\N	\N
45	2024-02-09 11:26:07.259216	afgafdga	\N	\N	\N	\N
50	2024-02-09 11:38:40.431663	great	\N	\N	\N	2003-10-12
\.


--
-- Data for Name: studentparentrelation; Type: TABLE DATA; Schema: school; Owner: vatsal
--

COPY school.studentparentrelation (student_id, parent_id) FROM stdin;
3	3
3	4
\.


--
-- Data for Name: subject; Type: TABLE DATA; Schema: school; Owner: vatsal
--

COPY school.subject (id, created_at, name) FROM stdin;
6	2024-02-07 15:04:35.024951	Physical Activity
7	2024-02-07 15:50:34.697209	Mathematics
8	2024-02-07 15:50:40.427391	Physics
\.


--
-- Data for Name: teacher; Type: TABLE DATA; Schema: school; Owner: vatsal
--

COPY school.teacher (id, created_at, name) FROM stdin;
3	2024-02-07 16:22:04.218067	DC SIR
4	2024-02-08 09:15:58.59096	Danaram Sir
5	2024-02-08 09:16:05.279478	Sonal Nair
6	2024-02-08 09:16:10.630875	Shilpa Yadav
7	2024-02-08 09:16:23.596043	Anil Mutange
\.


--
-- Data for Name: emp; Type: TABLE DATA; Schema: sql; Owner: vatsal
--

COPY sql.emp (id, name, age) FROM stdin;
4	Aman	45
5	Aman	48
6	Ketan	48
9	Ketan	48
10	Ketan	48
11	Ketan	48
12	Ketan	48
13	Ketan	48
14	Ketan	48
1	Aman	0
2	Aman	0
3	Aman	0
15	Ketan	48
16	Ketan	48
27	Ketan	48
28	Ketan	48
29	Ketan	48
30	Ketan	48
31	Ketan	48
32	Ketan	48
33	ketan	29
34	Ketan	89
35	ketan	29
\.


--
-- Data for Name: books; Type: TABLE DATA; Schema: task6muxgorm; Owner: vatsal
--

COPY task6muxgorm.books (id, title, author, isbn, publisher, year, genre, quantity) FROM stdin;
3	book	no one	abcdddffsf	Book vale	0	motivational	5
5	book	no one	abcdddffxdsf	Book vale	0	motivational	4
7	book	no one	abcdddffxerdsf	Book vale	0	motivational	5
9	book	no one	abcdddfwfxerdsf	Book vale	0	motivational	5
12	book	no one	abcdddfwefxerdsfd	Book vale	0	motivational	6
10	book	no one	abcdddfwefxerdsf	Book vale	0	motivational	5
14	book	no one	abcdddfwefxerdsfdf	Book vale	0	motivational	6
18	great	no one	UNIQUE	someone	0	Action	1
26	Test Book	Test Author	1234567890123		0		0
28	Test Book	Test Author	12345678901231		0		0
32	Test Book	Test Author	123456789012312		0		0
34	Test Book	Test Author	1234567389012312		0		0
43	testing book	someone	abcd		0		0
1	good	mice			0		0
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: task6muxgorm; Owner: vatsal
--

COPY task6muxgorm.users (id, name, password, role) FROM stdin;
1	ketan	123	admin
2	tridip 3334	222	user
\.


--
-- Data for Name: usersbooks; Type: TABLE DATA; Schema: task6muxgorm; Owner: vatsal
--

COPY task6muxgorm.usersbooks (user_id, book_id) FROM stdin;
1	3
1	1
1	5
2	5
2	7
2	9
2	10
\.


--
-- Data for Name: todos; Type: TABLE DATA; Schema: todo_cli; Owner: vatsal
--

COPY todo_cli.todos (id, title, completed, createdat) FROM stdin;
3	amazing-work	f	2024-02-08
1	wow	t	2024-02-08
2	wow	t	2024-02-08
4	do-work-on-sql	f	2024-02-08
5	newtodo	f	2024-02-08
6	newTodohaha	t	2024-02-08
7	great	t	2024-02-09
8	great-with-the-long-name-and-large-kingdome-hahahaha	f	2024-02-09
\.


--
-- Data for Name: bus; Type: TABLE DATA; Schema: transport; Owner: vatsal
--

COPY transport.bus (id, registration_number, model, capacity) FROM stdin;
2	GJ01XB8464	Ashok Leyland	50
3	GJ01XB6554	Mahindra	48
4	GJ01XB9808	TATA	50
5	GJ01XB5835	Ashok Leyland	45
6	GJ01XB5890	Mahindra	60
7	GJ01XB3536	TATA	50
8	GJ01XB1016	Ashok Leyland	70
9	GJ01XB3299	Mahindra	70
10	GJ01XB3246	TATA	70
11	GJ01XB2087	Ashok Leyland	45
12	GJ01XB7661	Mahindra	45
13	GJ01XB1784	TATA	60
14	GJ01XB6485	Eicher	60
15	GJ01XB5273	Eicher	60
16	GJ01XB8336	Ashok Leyland	50
17	GJ01XB8429	Mahindra	50
18	GJ01XB4248	TATA	50
19	GJ01XB5525	Ashok Leyland	45
20	GJ01XB3472	Mahindra	60
21	GJ01XB6827	TATA	50
22	GJ01XB5473	Ashok Leyland	70
23	GJ01XB9233	Mahindra	70
24	GJ01XB3507	TATA	70
25	GJ01XB5079	Ashok Leyland	45
26	GJ01XB2646	Mahindra	45
27	GJ01XB5602	TATA	60
28	GJ01XB7367	Eicher	60
29	GJ01XB3119	Eicher	60
30	GJ01XB7644	Ashok Leyland	50
31	GJ01XB2019	Mahindra	50
32	GJ01XB2718	TATA	50
33	GJ01XB9974	Ashok Leyland	45
34	GJ01XB9290	Mahindra	60
35	GJ01XB2985	TATA	50
36	GJ01XB3015	Ashok Leyland	70
37	GJ01XB4494	Mahindra	70
38	GJ01XB4914	TATA	70
39	GJ01XB5810	Ashok Leyland	45
40	GJ01XB5446	Mahindra	45
41	GJ01XB4712	TATA	60
42	GJ01XB1458	Eicher	60
1	GJ01XD5112	Volvo	20
57	GJ18AB1229	Mahindra	40
58	GJ01SB1536	Eicher	34
59	GJ18AB1005	Eicher	56
60	GJ18AB1034	Eicher	54
61	GJ18AB1010	Eicher	32
62	GJ18AB3524	Eicher	54
63	GJ18AB1967	Eicher	40
\.


--
-- Data for Name: busstatus; Type: TABLE DATA; Schema: transport; Owner: vatsal
--

COPY transport.busstatus (bus_id, lat, long, last_updated, traffic, status, last_station_order) FROM stdin;
18	0	0	00:00:00	0	0	0
19	0	0	00:00:00	0	0	0
20	0	0	00:00:00	0	0	0
21	0	0	00:00:00	0	0	0
22	0	0	00:00:00	0	0	0
23	0	0	00:00:00	0	0	0
24	0	0	00:00:00	0	0	0
25	0	0	00:00:00	0	0	0
26	0	0	00:00:00	0	0	0
27	0	0	00:00:00	0	0	0
28	0	0	00:00:00	0	0	0
29	0	0	00:00:00	0	0	0
30	0	0	00:00:00	0	0	0
31	0	0	00:00:00	0	0	0
32	0	0	00:00:00	0	0	0
2	23.0736327	72.592172	17:44:00	1	1	13
33	0	0	00:00:00	0	0	0
34	0	0	00:00:00	0	0	0
62	0	0	00:00:00	0	0	0
63	0	0	00:00:00	0	0	0
3	0	0	00:00:00	0	0	0
4	0	0	00:00:00	0	0	0
5	0	0	00:00:00	0	0	0
6	0	0	00:00:00	0	0	0
7	0	0	00:00:00	0	0	0
8	0	0	00:00:00	0	0	0
9	0	0	00:00:00	0	0	0
10	0	0	00:00:00	0	0	0
11	0	0	00:00:00	0	0	0
12	0	0	00:00:00	0	0	0
13	0	0	00:00:00	0	0	0
14	0	0	00:00:00	0	0	0
15	0	0	00:00:00	0	0	0
16	0	0	00:00:00	0	0	0
17	0	0	00:00:00	0	0	0
35	0	0	00:00:00	0	0	0
36	0	0	00:00:00	0	0	0
37	0	0	00:00:00	0	0	0
38	0	0	00:00:00	0	0	0
39	0	0	00:00:00	0	0	0
40	0	0	00:00:00	0	0	0
41	0	0	00:00:00	0	0	0
42	0	0	00:00:00	0	0	0
57	0	0	00:00:00	0	0	0
58	0	0	00:00:00	0	0	0
59	0	0	00:00:00	0	0	0
60	0	0	00:00:00	0	0	0
61	0	0	00:00:00	0	0	0
1	22.9768249	72.5841651	16:06:00	0	0	40
\.


--
-- Data for Name: driver; Type: TABLE DATA; Schema: transport; Owner: vatsal
--

COPY transport.driver (id, name, phone, gender, dob) FROM stdin;
1	Tridip	9290292323	0	2002-01-02
2	Tridip 2	9290292321	0	2002-01-01
3	Tridip 3	9290292322	0	2002-01-17
4	Tridip 4	9290292323	0	2002-01-02
5	Tridip 5	9290292323	0	2002-01-22
6	Tridip 6	9290292323	0	2002-01-02
7	Tridip 7	9290292323	0	2002-01-02
\.


--
-- Data for Name: route; Type: TABLE DATA; Schema: transport; Owner: vatsal
--

COPY transport.route (id, name, status, source, destination) FROM stdin;
1	7U	1	30	23
2	7D	1	23	30
3	7S	1	12	23
4	7S	1	24	23
5	7E	1	23	12
6	7E	1	23	24
7	4U	1	30	86
8	4D	1	86	30
9	6U	1	23	1
10	6D	1	1	23
11	6S	1	11	1
12	6S	1	24	1
13	6E	1	1	24
14	6E	1	1	11
15	8U	1	88	110
16	8D	1	110	88
\.


--
-- Data for Name: routestations; Type: TABLE DATA; Schema: transport; Owner: vatsal
--

COPY transport.routestations (route_id, station_id, station_order) FROM stdin;
10	1	1
10	2	2
10	3	3
10	4	4
10	5	5
10	6	6
10	7	7
10	8	8
10	9	9
10	10	10
10	11	11
10	12	12
10	13	13
10	14	14
10	15	15
10	16	16
10	17	17
10	18	18
10	19	19
10	20	20
10	21	21
10	22	22
10	23	23
9	23	1
9	22	2
9	21	3
9	20	4
9	19	5
9	18	6
9	17	7
9	16	8
9	15	9
9	14	10
9	13	11
9	12	12
9	11	13
9	10	14
9	9	15
9	8	16
9	7	17
9	6	18
9	5	19
9	4	20
9	3	21
9	2	22
9	1	23
11	11	1
11	10	2
11	9	3
11	8	4
11	7	5
11	6	6
11	5	7
11	4	8
11	3	9
11	2	10
11	1	11
1	30	1
1	31	2
1	32	3
1	33	4
1	34	5
1	35	6
1	36	7
1	37	8
1	38	9
1	39	10
1	40	11
1	41	12
1	42	13
1	43	14
1	44	15
1	45	16
1	46	17
1	47	18
1	48	19
1	49	20
1	50	21
1	51	22
1	52	23
1	53	24
1	54	25
1	55	26
1	56	27
1	57	28
1	58	29
1	59	30
1	60	31
1	61	32
1	62	33
1	63	34
1	64	35
1	65	36
1	66	37
1	67	38
1	68	39
1	69	40
1	23	41
2	23	1
2	69	2
2	68	3
2	67	4
2	66	5
2	65	6
2	64	7
2	63	8
2	62	9
2	61	10
2	60	11
2	59	12
2	58	13
2	57	14
2	56	15
2	55	16
2	54	17
2	53	18
2	52	19
2	51	20
2	50	21
2	49	22
2	48	23
2	47	24
2	46	25
2	45	26
2	44	27
2	43	28
2	42	29
2	41	30
2	40	31
2	39	32
2	38	33
2	37	34
2	36	35
2	35	36
2	34	37
2	33	38
2	32	39
2	31	40
2	30	41
3	12	1
3	13	2
3	14	3
3	15	4
3	16	5
3	17	6
3	18	7
3	19	8
3	20	9
3	21	10
3	22	11
3	23	12
5	23	1
5	22	2
5	21	3
5	20	4
5	19	5
5	18	6
5	17	7
5	16	8
5	15	9
5	14	10
5	13	11
5	12	12
4	24	1
4	25	2
4	26	3
4	27	4
4	28	5
4	29	6
4	11	7
4	12	8
4	13	9
4	14	10
4	15	11
4	16	12
4	17	13
4	18	14
4	19	15
4	20	16
4	21	17
4	22	18
4	23	19
6	23	1
6	22	2
6	21	3
6	20	4
6	19	5
6	18	6
6	17	7
6	16	8
6	15	9
6	14	10
6	13	11
6	12	12
6	11	13
6	29	14
6	28	15
6	27	16
6	26	17
6	25	18
6	24	19
7	30	1
7	31	2
7	32	3
7	33	4
7	34	5
7	35	6
7	37	7
7	38	8
7	39	9
7	40	10
7	41	11
7	42	12
7	43	13
7	44	14
7	70	15
7	71	16
7	72	17
7	73	18
7	74	19
7	75	20
7	76	21
7	77	22
7	78	23
7	79	24
7	80	25
7	81	26
7	82	27
7	83	28
7	84	29
7	85	30
7	86	31
7	87	32
8	87	1
8	86	2
8	85	3
8	84	4
8	83	5
8	82	6
8	81	7
8	80	8
8	79	9
8	78	10
8	77	11
8	76	12
8	75	13
8	74	14
8	73	15
8	72	16
8	71	17
8	70	18
8	44	19
8	43	20
8	42	21
8	41	22
8	40	23
8	39	24
8	38	25
8	37	26
8	35	27
8	34	28
8	33	29
8	32	30
8	31	31
8	30	32
12	24	1
12	25	2
12	26	3
12	27	4
12	28	5
12	29	6
12	9	7
12	8	8
12	7	9
12	6	10
12	5	11
12	4	12
12	3	13
12	2	14
12	1	15
13	1	1
13	2	2
13	3	3
13	4	4
13	5	5
13	6	6
13	7	7
13	8	8
13	9	9
13	29	10
13	28	11
13	27	12
13	26	13
13	25	14
13	24	15
14	1	1
14	2	2
14	3	3
14	4	4
14	5	5
14	6	6
14	7	7
14	8	8
14	9	9
14	10	10
14	11	11
15	88	1
15	89	2
15	90	3
15	91	4
15	92	5
15	81	6
15	82	7
15	83	8
15	84	9
15	93	10
15	86	11
15	87	12
15	94	13
15	95	14
15	96	15
15	97	16
15	98	17
15	99	18
15	100	19
15	59	20
15	58	21
15	57	22
15	56	23
15	55	24
15	101	25
15	102	26
15	103	27
15	104	28
15	105	29
15	106	30
15	107	31
15	108	32
15	1	33
15	109	34
15	110	35
16	110	1
16	109	2
16	1	3
16	108	4
16	107	5
16	106	6
16	105	7
16	104	8
16	103	9
16	102	10
16	101	11
16	54	12
16	55	13
16	56	14
16	57	15
16	58	16
16	59	17
16	100	18
16	99	19
16	98	20
16	97	21
16	96	22
16	95	23
16	94	24
16	87	25
16	86	26
16	93	27
16	84	28
16	83	29
16	82	30
16	81	31
16	92	32
16	91	33
16	90	34
16	89	35
16	88	36
\.


--
-- Data for Name: schedule; Type: TABLE DATA; Schema: transport; Owner: vatsal
--

COPY transport.schedule (id, bus_id, route_id, departure_time) FROM stdin;
1	1	1	07:30:00
2	2	1	11:30:00
3	3	1	15:30:00
4	4	2	07:30:00
5	5	2	11:30:00
6	6	2	15:30:00
7	7	3	08:00:00
8	8	3	10:00:00
9	9	3	14:00:00
10	10	4	09:00:00
11	11	4	11:00:00
12	12	4	18:00:00
13	13	5	08:00:00
14	14	5	10:00:00
15	15	5	14:00:00
16	16	6	09:00:00
17	17	6	11:00:00
18	18	6	18:00:00
19	19	7	08:00:00
20	20	7	12:00:00
21	21	7	16:00:00
22	22	8	08:00:00
23	23	8	12:00:00
24	24	8	16:00:00
25	25	9	10:00:00
26	26	9	15:00:00
27	27	9	18:00:00
28	28	10	10:00:00
29	29	10	15:00:00
30	30	10	18:00:00
31	31	11	06:30:00
32	32	11	10:30:00
33	33	11	13:00:00
34	34	12	07:30:00
35	35	12	11:30:00
36	36	12	14:00:00
37	37	13	06:30:00
38	38	13	10:30:00
39	39	13	13:00:00
40	40	14	07:30:00
41	41	14	11:30:00
42	42	14	14:00:00
43	57	15	07:30:00
44	58	15	11:30:00
45	59	15	17:30:00
46	60	16	07:30:00
47	61	16	12:30:00
48	62	16	15:30:00
\.


--
-- Data for Name: station; Type: TABLE DATA; Schema: transport; Owner: vatsal
--

COPY transport.station (id, name, lat, long) FROM stdin;
92	Jodhpur Char Rasta	23.05251	72.5270853
93	Panjrapole Char Rasta	23.026398	72.5434296
1	Naroda S. T. Workshop	23.0642697	72.6412547
2	Dhanus Dhari Mandir	23.059817	72.6403651
3	Krishnanagar	23.0538533	72.6412726
4	Vijay Park	23.051455	72.6412091
5	Hirawadi	23.0456452	72.640533
6	Thakkar Nagar Approach	23.0413413	72.6402739
7	Lilanagar	23.03724	72.6398149
8	Bapu Nagar Approach	23.0324414	72.6388862
9	Viratnagar	23.028364	72.6389126
10	Soni ni chali	23.0206933	72.6380236
11	Geeta Gauri Cinema	23.01732	72.6374156
12	Rameshwar Park	23.0132204	72.6377667
13	Ram Rajya Nagar	23.0079862	72.6373921
14	Rabari Colony	23.0044017	72.6359452
15	Jogeshwari Society	23.000486	72.6347704
16	Purvdeep Society	22.9959085	72.6331677
17	CTM Cross Road	22.9915924	72.6316459
18	Express Highway Junction	23.0182647	72.5530951
19	Jashodanagar Char Rasta	22.9838528	72.6214001
20	Ghodasar	22.9786829	72.6075398
21	Isanpur	22.9762188	72.5997748
22	Mukesh Industries	22.9740123	72.5926902
23	narol	22.9729729	72.5846005
24	Odhav Talav	23.0245023	72.6627271
25	Morlidhar Society	23.0246285	72.6592107
26	Chhotalal ni Chali	23.0235235	72.6561457
27	Vallabh Nagar	23.022949	72.6522757
28	Odhav Fire Station	23.0492335	72.5559743
29	Grid Station	23.0213	72.6409997
30	DCIS Circle	23.129273	72.5819126
31	Sarthi Bunglows	23.1229731	72.5778669
32	Chandkheda Gam	23.1165397	72.5819121
33	Shiv Shaktinagar	23.1125755	72.5838263
34	Jantangar	23.1097894	72.5844511
36	Vishvkarma Engineering	23.1058858	72.5929558
37	Visat – Gandhinagar Juncton	23.0974089	72.5859741
38	Motera Cross Road	23.0922298	72.5874837
39	Sabarmati Police Station	23.0861116	72.5919437
40	Sabarmati Municiple Swimming Pool	23.0815749	72.5907851
41	Rathi Apartment	23.0792364	72.5947427
42	Savarmati Power House	23.0736327	72.592172
43	R.T.O. Circle	23.0686563	72.581138
44	Ranip Cross Road	23.0675323	72.5749533
45	N R Patel Park	23.0665475	72.5714796
46	Ramapir No Tekro	23.061705	72.5718466
47	Juna Vadaj	23.0416688	72.5851244
48	Gurudwara	23.0476275	72.5779406
49	Hanumanpura	23.0449717	72.5787598
50	Sarkari Litho Press	23.0417749	72.5826251
51	Sarkari Litho Press Cabin	23.0394466	72.58619
52	Delhi Darwaja	23.0381455	72.5900518
53	Prem Darwaja	23.0341211	72.5962474
54	Kalupur Ghee Bazar	23.0311294	72.5970521
55	Kalupur	23.0281333	72.5974933
56	Sarangpur Darwaja	23.0217549	72.5964651
57	Karnamukteshwar Mahadev	23.0190749	72.5940491
58	Raipur Darwaja	23.0174141	72.5933785
59	Astodiya Darwaja	23.0168585	72.5896383
60	Geeta Mandir	23.0131047	72.5865086
61	Bhulabhai Park	23.0066929	72.5862481
62	Managal Park	23.0035763	72.5850253
63	Swaminarayan College	23.0007091	72.5723431
64	Vaikunth Dham Mandir	23.0006848	72.580068
65	Danilimda Road	22.9944244	72.575933
66	Chhipa Society	22.9897849	72.5760195
67	Chandola Lake	22.9872606	72.5824333
68	BRTS Workshop	22.9821729	72.5808051
69	Kashiram Textiles	22.9768249	72.5841651
70	Bhavsar Hostel	23.0676249	72.5660351
72	Pragatinagar	23.0646649	72.5511891
73	Shastrinagar	23.0622449	72.5504051
74	Jaimangal	23.0576179	72.5467631
75	Sola Cross Road	23.0527202	72.5419964
76	Shree Valinath Chowk	23.04912	72.5424251
77	Memnagar	23.0456096	72.5398606
78	University	23.0384082	72.5353596
79	Andhjan Mandal	23.0346443	72.5311408
80	Himmatlal Park	23.0300596	72.5300452
81	Shivranjani	23.0245279	72.5280903
82	Jhansi ki Rani	23.0230449	72.5337751
83	Nehrunagar	23.0229943	72.5331544
84	L Colony	23.0235749	72.5380091
85	Ranjrapole Char Rasta	23.0264029	72.5415681
86	Gulabai Tekra Approach	23.0294049	72.5446151
87	LD Engg. College	23.0349669	72.5462891
94	Vasundra Society	23.0277422	72.55205
35	ONGC Avani Bhavan	23.10357	72.5867515
88	Iskon Cross Road	23.027082	72.5082556
89	RamdevNagar	23.027082	72.5125358
90	ISRO	23.0276993	72.5174121
91	Star Bazaar	23.0260113	72.523829
95	Law Garden	23.0238385	72.558674
96	M J Library	23.023727	72.5699896
97	Lokmanya Tilak Bagh	23.0220613	72.5800264
98	Raikhad Char Rasta	23.0216164	72.5823638
99	Municipal Corporation Office	23.0197301	72.5859025
100	Astodiya Chakala	23.0187083	72.5883122
101	G.C.S Hospital	23.0390754	72.608098
102	Arvind Mill	23.0417658	72.6123563
103	Jeening Press	23.0443101	72.6154414
104	Ashok Mill	23.0465601	72.6185914
105	Naroda Fruit Market	23.0505771	72.6252054
106	Memco Cross Road	23.0533239	72.6291526
107	Muncipal North Zone Office	23.0557079	72.6328716
108	Saijpur Towers	23.0612901	72.6395926
109	Bethak	23.071791	72.6483232
110	Naroda Gam	23.0768739	72.6557968
71	Akhabar Nagar	23.066272	72.55878
\.


--
-- Name: schools_id_seq; Type: SEQUENCE SET; Schema: arrays; Owner: vatsal
--

SELECT pg_catalog.setval('arrays.schools_id_seq', 4, true);


--
-- Name: students_id_seq; Type: SEQUENCE SET; Schema: gormbasics; Owner: vatsal
--

SELECT pg_catalog.setval('gormbasics.students_id_seq', 76, true);


--
-- Name: creditcard_id_seq; Type: SEQUENCE SET; Schema: gormbasics1; Owner: vatsal
--

SELECT pg_catalog.setval('gormbasics1.creditcard_id_seq', 1, false);


--
-- Name: note_id_seq; Type: SEQUENCE SET; Schema: gormbasics1; Owner: vatsal
--

SELECT pg_catalog.setval('gormbasics1.note_id_seq', 3, true);


--
-- Name: product_id_seq; Type: SEQUENCE SET; Schema: gormbasics1; Owner: vatsal
--

SELECT pg_catalog.setval('gormbasics1.product_id_seq', 1, false);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: gormbasics1; Owner: vatsal
--

SELECT pg_catalog.setval('gormbasics1.users_id_seq', 1, false);


--
-- Name: school_id_seq; Type: SEQUENCE SET; Schema: gormschool; Owner: vatsal
--

SELECT pg_catalog.setval('gormschool.school_id_seq', 2, true);


--
-- Name: student_id_seq; Type: SEQUENCE SET; Schema: gormschool; Owner: vatsal
--

SELECT pg_catalog.setval('gormschool.student_id_seq', 5, true);


--
-- Name: parent_id_seq; Type: SEQUENCE SET; Schema: gormschoolproject; Owner: vatsal
--

SELECT pg_catalog.setval('gormschoolproject.parent_id_seq', 3, true);


--
-- Name: sport_id_seq; Type: SEQUENCE SET; Schema: gormschoolproject; Owner: vatsal
--

SELECT pg_catalog.setval('gormschoolproject.sport_id_seq', 5, true);


--
-- Name: student_id_seq; Type: SEQUENCE SET; Schema: gormschoolproject; Owner: vatsal
--

SELECT pg_catalog.setval('gormschoolproject.student_id_seq', 2345, true);


--
-- Name: user_id_seq; Type: SEQUENCE SET; Schema: httpnet; Owner: vatsal
--

SELECT pg_catalog.setval('httpnet.user_id_seq', 27, true);


--
-- Name: product_groups_group_id_seq; Type: SEQUENCE SET; Schema: products; Owner: vatsal
--

SELECT pg_catalog.setval('products.product_groups_group_id_seq', 3, true);


--
-- Name: products_product_id_seq; Type: SEQUENCE SET; Schema: products; Owner: vatsal
--

SELECT pg_catalog.setval('products.products_product_id_seq', 11, true);


--
-- Name: customers_id_seq; Type: SEQUENCE SET; Schema: public; Owner: vatsal
--

SELECT pg_catalog.setval('public.customers_id_seq', 6, true);


--
-- Name: emp_id_seq; Type: SEQUENCE SET; Schema: public; Owner: vatsal
--

SELECT pg_catalog.setval('public.emp_id_seq', 1, false);


--
-- Name: product_groups_group_id_seq; Type: SEQUENCE SET; Schema: public; Owner: vatsal
--

SELECT pg_catalog.setval('public.product_groups_group_id_seq', 3, true);


--
-- Name: products_id_seq; Type: SEQUENCE SET; Schema: public; Owner: vatsal
--

SELECT pg_catalog.setval('public.products_id_seq', 12, true);


--
-- Name: products_product_id_seq; Type: SEQUENCE SET; Schema: public; Owner: vatsal
--

SELECT pg_catalog.setval('public.products_product_id_seq', 12, true);


--
-- Name: studs_id_seq; Type: SEQUENCE SET; Schema: public; Owner: vatsal
--

SELECT pg_catalog.setval('public.studs_id_seq', 3, true);


--
-- Name: class_id_seq; Type: SEQUENCE SET; Schema: school; Owner: vatsal
--

SELECT pg_catalog.setval('school.class_id_seq', 5, true);


--
-- Name: parent_id_seq; Type: SEQUENCE SET; Schema: school; Owner: vatsal
--

SELECT pg_catalog.setval('school.parent_id_seq', 6, true);


--
-- Name: student_id_seq; Type: SEQUENCE SET; Schema: school; Owner: vatsal
--

SELECT pg_catalog.setval('school.student_id_seq', 2, true);


--
-- Name: subject_id_seq; Type: SEQUENCE SET; Schema: school; Owner: vatsal
--

SELECT pg_catalog.setval('school.subject_id_seq', 8, true);


--
-- Name: teacher_id_seq; Type: SEQUENCE SET; Schema: school; Owner: vatsal
--

SELECT pg_catalog.setval('school.teacher_id_seq', 7, true);


--
-- Name: emp_id_seq; Type: SEQUENCE SET; Schema: sql; Owner: vatsal
--

SELECT pg_catalog.setval('sql.emp_id_seq', 36, true);


--
-- Name: books_id_seq; Type: SEQUENCE SET; Schema: task6muxgorm; Owner: vatsal
--

SELECT pg_catalog.setval('task6muxgorm.books_id_seq', 43, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: task6muxgorm; Owner: vatsal
--

SELECT pg_catalog.setval('task6muxgorm.users_id_seq', 2, true);


--
-- Name: todos_id_seq; Type: SEQUENCE SET; Schema: todo_cli; Owner: vatsal
--

SELECT pg_catalog.setval('todo_cli.todos_id_seq', 8, true);


--
-- Name: bus_id_seq; Type: SEQUENCE SET; Schema: transport; Owner: vatsal
--

SELECT pg_catalog.setval('transport.bus_id_seq', 67, true);


--
-- Name: driver_id_seq; Type: SEQUENCE SET; Schema: transport; Owner: vatsal
--

SELECT pg_catalog.setval('transport.driver_id_seq', 8, true);


--
-- Name: route_id_seq; Type: SEQUENCE SET; Schema: transport; Owner: vatsal
--

SELECT pg_catalog.setval('transport.route_id_seq', 86, true);


--
-- Name: schedule_id_seq; Type: SEQUENCE SET; Schema: transport; Owner: vatsal
--

SELECT pg_catalog.setval('transport.schedule_id_seq', 1, false);


--
-- Name: station_id_seq; Type: SEQUENCE SET; Schema: transport; Owner: vatsal
--

SELECT pg_catalog.setval('transport.station_id_seq', 1, false);


--
-- Name: deestinct deestinct_id_key; Type: CONSTRAINT; Schema: doc5; Owner: vatsal
--

ALTER TABLE ONLY doc5.deestinct
    ADD CONSTRAINT deestinct_id_key UNIQUE (id);


--
-- Name: students students_pkey; Type: CONSTRAINT; Schema: gormbasics; Owner: vatsal
--

ALTER TABLE ONLY gormbasics.students
    ADD CONSTRAINT students_pkey PRIMARY KEY (id);


--
-- Name: category category_pkey; Type: CONSTRAINT; Schema: gormbasics1; Owner: vatsal
--

ALTER TABLE ONLY gormbasics1.category
    ADD CONSTRAINT category_pkey PRIMARY KEY (name);


--
-- Name: creditcard creditcard_pkey; Type: CONSTRAINT; Schema: gormbasics1; Owner: vatsal
--

ALTER TABLE ONLY gormbasics1.creditcard
    ADD CONSTRAINT creditcard_pkey PRIMARY KEY (id);


--
-- Name: note note_pkey; Type: CONSTRAINT; Schema: gormbasics1; Owner: vatsal
--

ALTER TABLE ONLY gormbasics1.note
    ADD CONSTRAINT note_pkey PRIMARY KEY (id);


--
-- Name: product product_pkey; Type: CONSTRAINT; Schema: gormbasics1; Owner: vatsal
--

ALTER TABLE ONLY gormbasics1.product
    ADD CONSTRAINT product_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: gormbasics1; Owner: vatsal
--

ALTER TABLE ONLY gormbasics1.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: school school_pkey; Type: CONSTRAINT; Schema: gormschool; Owner: vatsal
--

ALTER TABLE ONLY gormschool.school
    ADD CONSTRAINT school_pkey PRIMARY KEY (id);


--
-- Name: student student_pkey; Type: CONSTRAINT; Schema: gormschool; Owner: vatsal
--

ALTER TABLE ONLY gormschool.student
    ADD CONSTRAINT student_pkey PRIMARY KEY (id);


--
-- Name: parent parent_pkey; Type: CONSTRAINT; Schema: gormschoolproject; Owner: vatsal
--

ALTER TABLE ONLY gormschoolproject.parent
    ADD CONSTRAINT parent_pkey PRIMARY KEY (id);


--
-- Name: sport sport_pkey; Type: CONSTRAINT; Schema: gormschoolproject; Owner: vatsal
--

ALTER TABLE ONLY gormschoolproject.sport
    ADD CONSTRAINT sport_pkey PRIMARY KEY (id);


--
-- Name: student_parent student_parent_pkey; Type: CONSTRAINT; Schema: gormschoolproject; Owner: vatsal
--

ALTER TABLE ONLY gormschoolproject.student_parent
    ADD CONSTRAINT student_parent_pkey PRIMARY KEY (parent_id, student_id);


--
-- Name: student student_pkey; Type: CONSTRAINT; Schema: gormschoolproject; Owner: vatsal
--

ALTER TABLE ONLY gormschoolproject.student
    ADD CONSTRAINT student_pkey PRIMARY KEY (id);


--
-- Name: user user_email_key; Type: CONSTRAINT; Schema: httpnet; Owner: vatsal
--

ALTER TABLE ONLY httpnet."user"
    ADD CONSTRAINT user_email_key UNIQUE (email);


--
-- Name: user user_phone_key; Type: CONSTRAINT; Schema: httpnet; Owner: vatsal
--

ALTER TABLE ONLY httpnet."user"
    ADD CONSTRAINT user_phone_key UNIQUE (phone);


--
-- Name: user user_pkey; Type: CONSTRAINT; Schema: httpnet; Owner: vatsal
--

ALTER TABLE ONLY httpnet."user"
    ADD CONSTRAINT user_pkey PRIMARY KEY (id);


--
-- Name: product_groups product_groups_pkey; Type: CONSTRAINT; Schema: products; Owner: vatsal
--

ALTER TABLE ONLY products.product_groups
    ADD CONSTRAINT product_groups_pkey PRIMARY KEY (group_id);


--
-- Name: products products_pkey; Type: CONSTRAINT; Schema: products; Owner: vatsal
--

ALTER TABLE ONLY products.products
    ADD CONSTRAINT products_pkey PRIMARY KEY (product_id);


--
-- Name: categories categories_pkey; Type: CONSTRAINT; Schema: public; Owner: vatsal
--

ALTER TABLE ONLY public.categories
    ADD CONSTRAINT categories_pkey PRIMARY KEY (name);


--
-- Name: customers customers_pkey; Type: CONSTRAINT; Schema: public; Owner: vatsal
--

ALTER TABLE ONLY public.customers
    ADD CONSTRAINT customers_pkey PRIMARY KEY (id);


--
-- Name: product_groups product_groups_pkey; Type: CONSTRAINT; Schema: public; Owner: vatsal
--

ALTER TABLE ONLY public.product_groups
    ADD CONSTRAINT product_groups_pkey PRIMARY KEY (group_id);


--
-- Name: products products_pkey; Type: CONSTRAINT; Schema: public; Owner: vatsal
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_pkey PRIMARY KEY (product_id);


--
-- Name: student_parent student_parent_pkey; Type: CONSTRAINT; Schema: public; Owner: vatsal
--

ALTER TABLE ONLY public.student_parent
    ADD CONSTRAINT student_parent_pkey PRIMARY KEY (parent_id, student_id);


--
-- Name: studs studs_pkey; Type: CONSTRAINT; Schema: public; Owner: vatsal
--

ALTER TABLE ONLY public.studs
    ADD CONSTRAINT studs_pkey PRIMARY KEY (id);


--
-- Name: class class_pkey; Type: CONSTRAINT; Schema: school; Owner: vatsal
--

ALTER TABLE ONLY school.class
    ADD CONSTRAINT class_pkey PRIMARY KEY (id);


--
-- Name: classsubjectrelation classsubjectrelation_pkey; Type: CONSTRAINT; Schema: school; Owner: vatsal
--

ALTER TABLE ONLY school.classsubjectrelation
    ADD CONSTRAINT classsubjectrelation_pkey PRIMARY KEY (class_id, subject_id);


--
-- Name: parent parent_pkey; Type: CONSTRAINT; Schema: school; Owner: vatsal
--

ALTER TABLE ONLY school.parent
    ADD CONSTRAINT parent_pkey PRIMARY KEY (id);


--
-- Name: registration pk_registration; Type: CONSTRAINT; Schema: school; Owner: vatsal
--

ALTER TABLE ONLY school.registration
    ADD CONSTRAINT pk_registration PRIMARY KEY (student_id, year);


--
-- Name: student student_pkey; Type: CONSTRAINT; Schema: school; Owner: vatsal
--

ALTER TABLE ONLY school.student
    ADD CONSTRAINT student_pkey PRIMARY KEY (id);


--
-- Name: studentparentrelation studentparentrelation_pkey; Type: CONSTRAINT; Schema: school; Owner: vatsal
--

ALTER TABLE ONLY school.studentparentrelation
    ADD CONSTRAINT studentparentrelation_pkey PRIMARY KEY (student_id, parent_id);


--
-- Name: subject subject_pkey; Type: CONSTRAINT; Schema: school; Owner: vatsal
--

ALTER TABLE ONLY school.subject
    ADD CONSTRAINT subject_pkey PRIMARY KEY (id);


--
-- Name: teacher teacher_pkey; Type: CONSTRAINT; Schema: school; Owner: vatsal
--

ALTER TABLE ONLY school.teacher
    ADD CONSTRAINT teacher_pkey PRIMARY KEY (id);


--
-- Name: books books_pkey; Type: CONSTRAINT; Schema: task6muxgorm; Owner: vatsal
--

ALTER TABLE ONLY task6muxgorm.books
    ADD CONSTRAINT books_pkey PRIMARY KEY (id);


--
-- Name: books uni_task6muxgorm_books_isbn; Type: CONSTRAINT; Schema: task6muxgorm; Owner: vatsal
--

ALTER TABLE ONLY task6muxgorm.books
    ADD CONSTRAINT uni_task6muxgorm_books_isbn UNIQUE (isbn);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: task6muxgorm; Owner: vatsal
--

ALTER TABLE ONLY task6muxgorm.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: usersbooks usersbooks_pkey; Type: CONSTRAINT; Schema: task6muxgorm; Owner: vatsal
--

ALTER TABLE ONLY task6muxgorm.usersbooks
    ADD CONSTRAINT usersbooks_pkey PRIMARY KEY (user_id, book_id);


--
-- Name: bus bus_pkey; Type: CONSTRAINT; Schema: transport; Owner: vatsal
--

ALTER TABLE ONLY transport.bus
    ADD CONSTRAINT bus_pkey PRIMARY KEY (id);


--
-- Name: bus bus_registration_number_key; Type: CONSTRAINT; Schema: transport; Owner: vatsal
--

ALTER TABLE ONLY transport.bus
    ADD CONSTRAINT bus_registration_number_key UNIQUE (registration_number);


--
-- Name: busstatus busstatus_bus_id_key; Type: CONSTRAINT; Schema: transport; Owner: vatsal
--

ALTER TABLE ONLY transport.busstatus
    ADD CONSTRAINT busstatus_bus_id_key UNIQUE (bus_id);


--
-- Name: driver driver_id_key; Type: CONSTRAINT; Schema: transport; Owner: vatsal
--

ALTER TABLE ONLY transport.driver
    ADD CONSTRAINT driver_id_key UNIQUE (id);


--
-- Name: driver driver_pkey; Type: CONSTRAINT; Schema: transport; Owner: vatsal
--

ALTER TABLE ONLY transport.driver
    ADD CONSTRAINT driver_pkey PRIMARY KEY (id);


--
-- Name: route route_pkey; Type: CONSTRAINT; Schema: transport; Owner: vatsal
--

ALTER TABLE ONLY transport.route
    ADD CONSTRAINT route_pkey PRIMARY KEY (id);


--
-- Name: routestations routestations_pkey; Type: CONSTRAINT; Schema: transport; Owner: vatsal
--

ALTER TABLE ONLY transport.routestations
    ADD CONSTRAINT routestations_pkey PRIMARY KEY (route_id, station_id, station_order);


--
-- Name: schedule schedule_pkey; Type: CONSTRAINT; Schema: transport; Owner: vatsal
--

ALTER TABLE ONLY transport.schedule
    ADD CONSTRAINT schedule_pkey PRIMARY KEY (id);


--
-- Name: station station_pkey; Type: CONSTRAINT; Schema: transport; Owner: vatsal
--

ALTER TABLE ONLY transport.station
    ADD CONSTRAINT station_pkey PRIMARY KEY (id);


--
-- Name: idx_gormbasics_students_deleted_at; Type: INDEX; Schema: gormbasics; Owner: vatsal
--

CREATE INDEX idx_gormbasics_students_deleted_at ON gormbasics.students USING btree (deleted_at);


--
-- Name: idx_gormbasics1_creditcard_deleted_at; Type: INDEX; Schema: gormbasics1; Owner: vatsal
--

CREATE INDEX idx_gormbasics1_creditcard_deleted_at ON gormbasics1.creditcard USING btree (deleted_at);


--
-- Name: idx_gormbasics1_note_deleted_at; Type: INDEX; Schema: gormbasics1; Owner: vatsal
--

CREATE INDEX idx_gormbasics1_note_deleted_at ON gormbasics1.note USING btree (deleted_at);


--
-- Name: idx_gormbasics1_note_user_id; Type: INDEX; Schema: gormbasics1; Owner: vatsal
--

CREATE INDEX idx_gormbasics1_note_user_id ON gormbasics1.note USING btree (user_id);


--
-- Name: idx_gormbasics1_users_deleted_at; Type: INDEX; Schema: gormbasics1; Owner: vatsal
--

CREATE INDEX idx_gormbasics1_users_deleted_at ON gormbasics1.users USING btree (deleted_at);


--
-- Name: idx_customers_deleted_at; Type: INDEX; Schema: public; Owner: vatsal
--

CREATE INDEX idx_customers_deleted_at ON public.customers USING btree (deleted_at);


--
-- Name: idx_products_deleted_at; Type: INDEX; Schema: public; Owner: vatsal
--

CREATE INDEX idx_products_deleted_at ON public.products USING btree (deleted_at);


--
-- Name: idx_studs_deleted_at; Type: INDEX; Schema: public; Owner: vatsal
--

CREATE INDEX idx_studs_deleted_at ON public.studs USING btree (deleted_at);


--
-- Name: student calculate_age_trigger; Type: TRIGGER; Schema: school; Owner: vatsal
--

CREATE TRIGGER calculate_age_trigger BEFORE INSERT ON school.student FOR EACH ROW EXECUTE FUNCTION public.calculate_age();


--
-- Name: student update_age; Type: TRIGGER; Schema: school; Owner: vatsal
--

CREATE TRIGGER update_age BEFORE INSERT ON school.student FOR EACH ROW EXECUTE FUNCTION public.calculate_age();


--
-- Name: routestations routeupdatetrigger; Type: TRIGGER; Schema: transport; Owner: vatsal
--

CREATE TRIGGER routeupdatetrigger BEFORE INSERT ON transport.routestations FOR EACH STATEMENT EXECUTE FUNCTION public.funcforroute();


--
-- Name: routestations triggerupdateroute; Type: TRIGGER; Schema: transport; Owner: vatsal
--

CREATE TRIGGER triggerupdateroute AFTER INSERT OR DELETE OR UPDATE ON transport.routestations FOR EACH ROW EXECUTE FUNCTION public.updateviewroute();


--
-- Name: schedule triggerupdateroute; Type: TRIGGER; Schema: transport; Owner: vatsal
--

CREATE TRIGGER triggerupdateroute AFTER INSERT OR DELETE OR UPDATE ON transport.schedule FOR EACH ROW EXECUTE FUNCTION public.updateviewroute();


--
-- Name: product fk_gormbasics1_category_products; Type: FK CONSTRAINT; Schema: gormbasics1; Owner: vatsal
--

ALTER TABLE ONLY gormbasics1.product
    ADD CONSTRAINT fk_gormbasics1_category_products FOREIGN KEY (category_name) REFERENCES gormbasics1.category(name);


--
-- Name: note fk_gormbasics1_note_user; Type: FK CONSTRAINT; Schema: gormbasics1; Owner: vatsal
--

ALTER TABLE ONLY gormbasics1.note
    ADD CONSTRAINT fk_gormbasics1_note_user FOREIGN KEY (user_id) REFERENCES gormbasics1.users(id);


--
-- Name: creditcard fk_gormbasics1_users_credit_card; Type: FK CONSTRAINT; Schema: gormbasics1; Owner: vatsal
--

ALTER TABLE ONLY gormbasics1.creditcard
    ADD CONSTRAINT fk_gormbasics1_users_credit_card FOREIGN KEY (user_id) REFERENCES gormbasics1.users(id);


--
-- Name: note fk_gormbasics1_users_notes; Type: FK CONSTRAINT; Schema: gormbasics1; Owner: vatsal
--

ALTER TABLE ONLY gormbasics1.note
    ADD CONSTRAINT fk_gormbasics1_users_notes FOREIGN KEY (user_id) REFERENCES gormbasics1.users(id);


--
-- Name: student fk_gormschool_student_school; Type: FK CONSTRAINT; Schema: gormschool; Owner: vatsal
--

ALTER TABLE ONLY gormschool.student
    ADD CONSTRAINT fk_gormschool_student_school FOREIGN KEY (school_id) REFERENCES gormschool.school(id);


--
-- Name: student_parent fk_gormschoolproject_student_parent_parent; Type: FK CONSTRAINT; Schema: gormschoolproject; Owner: vatsal
--

ALTER TABLE ONLY gormschoolproject.student_parent
    ADD CONSTRAINT fk_gormschoolproject_student_parent_parent FOREIGN KEY (parent_id) REFERENCES gormschoolproject.parent(id);


--
-- Name: student_parent fk_gormschoolproject_student_parent_student; Type: FK CONSTRAINT; Schema: gormschoolproject; Owner: vatsal
--

ALTER TABLE ONLY gormschoolproject.student_parent
    ADD CONSTRAINT fk_gormschoolproject_student_parent_student FOREIGN KEY (student_id) REFERENCES gormschoolproject.student(id);


--
-- Name: student fk_gormschoolproject_student_sport; Type: FK CONSTRAINT; Schema: gormschoolproject; Owner: vatsal
--

ALTER TABLE ONLY gormschoolproject.student
    ADD CONSTRAINT fk_gormschoolproject_student_sport FOREIGN KEY (sport_id) REFERENCES gormschoolproject.sport(id);


--
-- Name: products products_group_id_fkey; Type: FK CONSTRAINT; Schema: products; Owner: vatsal
--

ALTER TABLE ONLY products.products
    ADD CONSTRAINT products_group_id_fkey FOREIGN KEY (group_id) REFERENCES public.product_groups(group_id);


--
-- Name: student_parent fk_student_parent_parent; Type: FK CONSTRAINT; Schema: public; Owner: vatsal
--

ALTER TABLE ONLY public.student_parent
    ADD CONSTRAINT fk_student_parent_parent FOREIGN KEY (parent_id) REFERENCES gormschoolproject.parent(id);


--
-- Name: student_parent fk_student_parent_student; Type: FK CONSTRAINT; Schema: public; Owner: vatsal
--

ALTER TABLE ONLY public.student_parent
    ADD CONSTRAINT fk_student_parent_student FOREIGN KEY (student_id) REFERENCES gormschoolproject.student(id);


--
-- Name: products products_group_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: vatsal
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_group_id_fkey FOREIGN KEY (group_id) REFERENCES public.product_groups(group_id);


--
-- Name: classsubjectrelation classsubjectrelation_class_id_fkey; Type: FK CONSTRAINT; Schema: school; Owner: vatsal
--

ALTER TABLE ONLY school.classsubjectrelation
    ADD CONSTRAINT classsubjectrelation_class_id_fkey FOREIGN KEY (class_id) REFERENCES school.class(id);


--
-- Name: classsubjectrelation classsubjectrelation_subject_id_fkey; Type: FK CONSTRAINT; Schema: school; Owner: vatsal
--

ALTER TABLE ONLY school.classsubjectrelation
    ADD CONSTRAINT classsubjectrelation_subject_id_fkey FOREIGN KEY (subject_id) REFERENCES school.subject(id);


--
-- Name: classsubjectrelation fk_class_subject; Type: FK CONSTRAINT; Schema: school; Owner: vatsal
--

ALTER TABLE ONLY school.classsubjectrelation
    ADD CONSTRAINT fk_class_subject FOREIGN KEY (teacher_id) REFERENCES school.teacher(id) ON DELETE SET NULL;


--
-- Name: registration registration_class_id_fkey; Type: FK CONSTRAINT; Schema: school; Owner: vatsal
--

ALTER TABLE ONLY school.registration
    ADD CONSTRAINT registration_class_id_fkey FOREIGN KEY (class_id) REFERENCES school.class(id);


--
-- Name: registration registration_student_id_fkey; Type: FK CONSTRAINT; Schema: school; Owner: vatsal
--

ALTER TABLE ONLY school.registration
    ADD CONSTRAINT registration_student_id_fkey FOREIGN KEY (student_id) REFERENCES school.student(id);


--
-- Name: studentparentrelation studentparentrelation_relation_1; Type: FK CONSTRAINT; Schema: school; Owner: vatsal
--

ALTER TABLE ONLY school.studentparentrelation
    ADD CONSTRAINT studentparentrelation_relation_1 FOREIGN KEY (student_id) REFERENCES school.student(id) ON DELETE CASCADE;


--
-- Name: studentparentrelation studentparentrelation_relation_2; Type: FK CONSTRAINT; Schema: school; Owner: vatsal
--

ALTER TABLE ONLY school.studentparentrelation
    ADD CONSTRAINT studentparentrelation_relation_2 FOREIGN KEY (parent_id) REFERENCES school.parent(id) ON DELETE CASCADE;


--
-- Name: usersbooks fk_task6muxgorm_usersbooks_book; Type: FK CONSTRAINT; Schema: task6muxgorm; Owner: vatsal
--

ALTER TABLE ONLY task6muxgorm.usersbooks
    ADD CONSTRAINT fk_task6muxgorm_usersbooks_book FOREIGN KEY (book_id) REFERENCES task6muxgorm.books(id);


--
-- Name: usersbooks fk_task6muxgorm_usersbooks_user; Type: FK CONSTRAINT; Schema: task6muxgorm; Owner: vatsal
--

ALTER TABLE ONLY task6muxgorm.usersbooks
    ADD CONSTRAINT fk_task6muxgorm_usersbooks_user FOREIGN KEY (user_id) REFERENCES task6muxgorm.users(id);


--
-- Name: busstatus busstatus_bus_id_fkey; Type: FK CONSTRAINT; Schema: transport; Owner: vatsal
--

ALTER TABLE ONLY transport.busstatus
    ADD CONSTRAINT busstatus_bus_id_fkey FOREIGN KEY (bus_id) REFERENCES transport.bus(id);


--
-- Name: route route_destination_fkey; Type: FK CONSTRAINT; Schema: transport; Owner: vatsal
--

ALTER TABLE ONLY transport.route
    ADD CONSTRAINT route_destination_fkey FOREIGN KEY (destination) REFERENCES transport.station(id) ON DELETE RESTRICT;


--
-- Name: routestations route_id_fkey; Type: FK CONSTRAINT; Schema: transport; Owner: vatsal
--

ALTER TABLE ONLY transport.routestations
    ADD CONSTRAINT route_id_fkey FOREIGN KEY (route_id) REFERENCES transport.route(id) ON DELETE CASCADE;


--
-- Name: route route_source_fkey; Type: FK CONSTRAINT; Schema: transport; Owner: vatsal
--

ALTER TABLE ONLY transport.route
    ADD CONSTRAINT route_source_fkey FOREIGN KEY (source) REFERENCES transport.station(id) ON DELETE RESTRICT;


--
-- Name: schedule schedule_bus_id_fkey; Type: FK CONSTRAINT; Schema: transport; Owner: vatsal
--

ALTER TABLE ONLY transport.schedule
    ADD CONSTRAINT schedule_bus_id_fkey FOREIGN KEY (bus_id) REFERENCES transport.bus(id) ON DELETE CASCADE;


--
-- Name: schedule schedule_route_id_fkey; Type: FK CONSTRAINT; Schema: transport; Owner: vatsal
--

ALTER TABLE ONLY transport.schedule
    ADD CONSTRAINT schedule_route_id_fkey FOREIGN KEY (route_id) REFERENCES transport.route(id) ON DELETE CASCADE;


--
-- Name: routestations station_id_fkey; Type: FK CONSTRAINT; Schema: transport; Owner: vatsal
--

ALTER TABLE ONLY transport.routestations
    ADD CONSTRAINT station_id_fkey FOREIGN KEY (station_id) REFERENCES transport.station(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

