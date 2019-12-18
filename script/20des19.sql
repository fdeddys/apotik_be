--
-- PostgreSQL database dump
--

-- Dumped from database version 10.5
-- Dumped by pg_dump version 10.5

-- Started on 2019-12-18 20:16:18 WIB

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 1 (class 3079 OID 12544)
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- TOC entry 2704 (class 0 OID 0)
-- Dependencies: 1
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET default_tablespace = '';

SET default_with_oids = false;

--
-- TOC entry 244 (class 1259 OID 20528)
-- Name: adjustment; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.adjustment (
    id bigint NOT NULL,
    adjustment_no character varying(25),
    adjustment_date timestamp with time zone,
    note character varying(255),
    last_update_by character varying(100),
    last_update timestamp without time zone,
    total numeric,
    status integer DEFAULT 0
);


ALTER TABLE public.adjustment OWNER TO postgres;

--
-- TOC entry 246 (class 1259 OID 20540)
-- Name: adjustment_detail; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.adjustment_detail (
    id bigint NOT NULL,
    adjustment_id bigint,
    product_id bigint DEFAULT 0 NOT NULL,
    qty numeric DEFAULT 0,
    hpp numeric DEFAULT 0,
    uom bigint DEFAULT 0 NOT NULL,
    last_update_by character varying(100),
    last_update timestamp without time zone
);


ALTER TABLE public.adjustment_detail OWNER TO postgres;

--
-- TOC entry 245 (class 1259 OID 20538)
-- Name: adjustment_detail_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.adjustment_detail_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.adjustment_detail_id_seq OWNER TO postgres;

--
-- TOC entry 2705 (class 0 OID 0)
-- Dependencies: 245
-- Name: adjustment_detail_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.adjustment_detail_id_seq OWNED BY public.adjustment_detail.id;


--
-- TOC entry 243 (class 1259 OID 20526)
-- Name: adjustment_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.adjustment_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.adjustment_id_seq OWNER TO postgres;

--
-- TOC entry 2706 (class 0 OID 0)
-- Dependencies: 243
-- Name: adjustment_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.adjustment_id_seq OWNED BY public.adjustment.id;


--
-- TOC entry 208 (class 1259 OID 20160)
-- Name: brand; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.brand (
    id bigint NOT NULL,
    status bigint NOT NULL,
    code character varying(100),
    name character varying(250),
    last_update_by character varying(100),
    last_update timestamp without time zone
);


ALTER TABLE public.brand OWNER TO postgres;

--
-- TOC entry 207 (class 1259 OID 20158)
-- Name: brand_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.brand_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.brand_id_seq OWNER TO postgres;

--
-- TOC entry 2707 (class 0 OID 0)
-- Dependencies: 207
-- Name: brand_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.brand_id_seq OWNED BY public.brand.id;


--
-- TOC entry 197 (class 1259 OID 20111)
-- Name: customer; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.customer (
    id bigint NOT NULL,
    code character varying(10),
    name character varying(100),
    last_update_by character varying(100),
    last_update timestamp without time zone,
    top integer DEFAULT 0,
    status integer DEFAULT 1,
    address1 character varying(500),
    address2 character varying(500),
    address3 character varying(500),
    address4 character varying(500),
    contact_person character varying(100),
    phone_number character varying(100)
);


ALTER TABLE public.customer OWNER TO postgres;

--
-- TOC entry 196 (class 1259 OID 20109)
-- Name: customer_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.customer_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.customer_id_seq OWNER TO postgres;

--
-- TOC entry 2708 (class 0 OID 0)
-- Dependencies: 196
-- Name: customer_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.customer_id_seq OWNED BY public.customer.id;


--
-- TOC entry 238 (class 1259 OID 20349)
-- Name: history_stock; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.history_stock (
    id bigint NOT NULL,
    code character varying(100),
    name character varying(250),
    debet real DEFAULT 0 NOT NULL,
    kredit real DEFAULT 0 NOT NULL,
    saldo real DEFAULT 0 NOT NULL,
    trans_date timestamp without time zone,
    description character varying(500),
    last_update_by character varying(100),
    last_update timestamp without time zone,
    reff_no character varying(100),
    price real DEFAULT 0,
    hpp character varying DEFAULT 0
);


ALTER TABLE public.history_stock OWNER TO postgres;

--
-- TOC entry 235 (class 1259 OID 20343)
-- Name: history_stock_debet_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.history_stock_debet_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.history_stock_debet_seq OWNER TO postgres;

--
-- TOC entry 2709 (class 0 OID 0)
-- Dependencies: 235
-- Name: history_stock_debet_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.history_stock_debet_seq OWNED BY public.history_stock.debet;


--
-- TOC entry 234 (class 1259 OID 20341)
-- Name: history_stock_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.history_stock_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.history_stock_id_seq OWNER TO postgres;

--
-- TOC entry 2710 (class 0 OID 0)
-- Dependencies: 234
-- Name: history_stock_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.history_stock_id_seq OWNED BY public.history_stock.id;


--
-- TOC entry 236 (class 1259 OID 20345)
-- Name: history_stock_kredit_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.history_stock_kredit_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.history_stock_kredit_seq OWNER TO postgres;

--
-- TOC entry 2711 (class 0 OID 0)
-- Dependencies: 236
-- Name: history_stock_kredit_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.history_stock_kredit_seq OWNED BY public.history_stock.kredit;


--
-- TOC entry 237 (class 1259 OID 20347)
-- Name: history_stock_saldo_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.history_stock_saldo_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.history_stock_saldo_seq OWNER TO postgres;

--
-- TOC entry 2712 (class 0 OID 0)
-- Dependencies: 237
-- Name: history_stock_saldo_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.history_stock_saldo_seq OWNED BY public.history_stock.saldo;


--
-- TOC entry 201 (class 1259 OID 20132)
-- Name: lookup; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.lookup (
    id bigint NOT NULL,
    status bigint DEFAULT 1,
    code character varying(100),
    lookup_group character varying(255),
    name character varying(250),
    is_viewable integer DEFAULT 1
);


ALTER TABLE public.lookup OWNER TO postgres;

--
-- TOC entry 240 (class 1259 OID 20422)
-- Name: lookup_group; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.lookup_group (
    id bigint NOT NULL,
    name character varying(250)
);


ALTER TABLE public.lookup_group OWNER TO postgres;

--
-- TOC entry 239 (class 1259 OID 20420)
-- Name: lookup_group_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.lookup_group_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.lookup_group_id_seq OWNER TO postgres;

--
-- TOC entry 2713 (class 0 OID 0)
-- Dependencies: 239
-- Name: lookup_group_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.lookup_group_id_seq OWNED BY public.lookup_group.id;


--
-- TOC entry 200 (class 1259 OID 20130)
-- Name: lookup_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.lookup_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.lookup_id_seq OWNER TO postgres;

--
-- TOC entry 2714 (class 0 OID 0)
-- Dependencies: 200
-- Name: lookup_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.lookup_id_seq OWNED BY public.lookup.id;


--
-- TOC entry 210 (class 1259 OID 20168)
-- Name: m_menus; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.m_menus (
    id bigint NOT NULL,
    name character varying(30) NOT NULL,
    description character varying(100) NOT NULL,
    last_update_by character varying(100),
    last_update timestamp without time zone,
    link character varying(200),
    parent_id bigint,
    icon character varying(50),
    status integer
);


ALTER TABLE public.m_menus OWNER TO postgres;

--
-- TOC entry 209 (class 1259 OID 20166)
-- Name: m_menus_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.m_menus_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.m_menus_id_seq OWNER TO postgres;

--
-- TOC entry 2715 (class 0 OID 0)
-- Dependencies: 209
-- Name: m_menus_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.m_menus_id_seq OWNED BY public.m_menus.id;


--
-- TOC entry 213 (class 1259 OID 20182)
-- Name: m_role_menu; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.m_role_menu (
    role_id bigint NOT NULL,
    menu_id bigint NOT NULL,
    status integer,
    last_update_by character varying(100),
    last_update timestamp without time zone
);


ALTER TABLE public.m_role_menu OWNER TO postgres;

--
-- TOC entry 216 (class 1259 OID 20198)
-- Name: m_role_user; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.m_role_user (
    role_id bigint NOT NULL,
    user_id bigint NOT NULL,
    last_update_by character varying(100),
    last_update timestamp without time zone,
    status integer DEFAULT 0
);


ALTER TABLE public.m_role_user OWNER TO postgres;

--
-- TOC entry 212 (class 1259 OID 20176)
-- Name: m_roles; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.m_roles (
    id bigint NOT NULL,
    name character varying(50) NOT NULL,
    description character varying(255) NOT NULL,
    last_update_by character varying(100),
    last_update timestamp without time zone
);


ALTER TABLE public.m_roles OWNER TO postgres;

--
-- TOC entry 211 (class 1259 OID 20174)
-- Name: m_roles_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.m_roles_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.m_roles_id_seq OWNER TO postgres;

--
-- TOC entry 2716 (class 0 OID 0)
-- Dependencies: 211
-- Name: m_roles_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.m_roles_id_seq OWNED BY public.m_roles.id;


--
-- TOC entry 215 (class 1259 OID 20187)
-- Name: m_users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.m_users (
    id bigint NOT NULL,
    user_name character varying(50) NOT NULL,
    email character varying(50) NOT NULL,
    password character varying(100) NOT NULL,
    status integer NOT NULL,
    last_update_by character varying(100),
    last_update timestamp without time zone,
    first_name character varying(50) DEFAULT '-'::character varying NOT NULL,
    last_name character varying(50) DEFAULT '-'::character varying NOT NULL,
    role_id bigint
);


ALTER TABLE public.m_users OWNER TO postgres;

--
-- TOC entry 214 (class 1259 OID 20185)
-- Name: m_users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.m_users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.m_users_id_seq OWNER TO postgres;

--
-- TOC entry 2717 (class 0 OID 0)
-- Dependencies: 214
-- Name: m_users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.m_users_id_seq OWNED BY public.m_users.id;


--
-- TOC entry 222 (class 1259 OID 20222)
-- Name: sales_order; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sales_order (
    id bigint NOT NULL,
    sales_order_no character varying(25),
    order_date timestamp with time zone,
    customer_id bigint,
    note character varying(255),
    last_update_by character varying(100),
    last_update timestamp without time zone,
    tax real,
    total real,
    grand_total real,
    salesman_id bigint DEFAULT 0,
    status integer DEFAULT 0,
    top integer DEFAULT 0,
    is_cash integer DEFAULT 0
);


ALTER TABLE public.sales_order OWNER TO postgres;

--
-- TOC entry 220 (class 1259 OID 20218)
-- Name: order_customer_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.order_customer_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.order_customer_id_seq OWNER TO postgres;

--
-- TOC entry 2718 (class 0 OID 0)
-- Dependencies: 220
-- Name: order_customer_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.order_customer_id_seq OWNED BY public.sales_order.customer_id;


--
-- TOC entry 219 (class 1259 OID 20216)
-- Name: order_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.order_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.order_id_seq OWNER TO postgres;

--
-- TOC entry 2719 (class 0 OID 0)
-- Dependencies: 219
-- Name: order_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.order_id_seq OWNED BY public.sales_order.id;


--
-- TOC entry 221 (class 1259 OID 20220)
-- Name: order_salesman_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.order_salesman_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.order_salesman_id_seq OWNER TO postgres;

--
-- TOC entry 2720 (class 0 OID 0)
-- Dependencies: 221
-- Name: order_salesman_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.order_salesman_id_seq OWNED BY public.sales_order.salesman_id;


--
-- TOC entry 206 (class 1259 OID 20149)
-- Name: product; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.product (
    id bigint NOT NULL,
    code character varying(10),
    name character varying(255),
    product_group_id bigint,
    brand_id bigint,
    small_uom_id bigint,
    status integer,
    last_update_by character varying(100),
    last_update timestamp without time zone,
    big_uom_id bigint,
    qty_uom integer DEFAULT 1,
    qty_stock real DEFAULT 0,
    hpp real DEFAULT 0,
    sell_price real
);


ALTER TABLE public.product OWNER TO postgres;

--
-- TOC entry 204 (class 1259 OID 20145)
-- Name: product_brand_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.product_brand_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.product_brand_id_seq OWNER TO postgres;

--
-- TOC entry 2721 (class 0 OID 0)
-- Dependencies: 204
-- Name: product_brand_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.product_brand_id_seq OWNED BY public.product.brand_id;


--
-- TOC entry 199 (class 1259 OID 20124)
-- Name: product_group; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.product_group (
    id bigint NOT NULL,
    name character varying(255),
    status integer DEFAULT 1,
    last_update_by character varying(100),
    last_update timestamp without time zone,
    code character varying(10)
);


ALTER TABLE public.product_group OWNER TO postgres;

--
-- TOC entry 198 (class 1259 OID 20122)
-- Name: product_group_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.product_group_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.product_group_id_seq OWNER TO postgres;

--
-- TOC entry 2722 (class 0 OID 0)
-- Dependencies: 198
-- Name: product_group_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.product_group_id_seq OWNED BY public.product_group.id;


--
-- TOC entry 202 (class 1259 OID 20141)
-- Name: product_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.product_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.product_id_seq OWNER TO postgres;

--
-- TOC entry 2723 (class 0 OID 0)
-- Dependencies: 202
-- Name: product_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.product_id_seq OWNED BY public.product.id;


--
-- TOC entry 203 (class 1259 OID 20143)
-- Name: product_product_group_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.product_product_group_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.product_product_group_id_seq OWNER TO postgres;

--
-- TOC entry 2724 (class 0 OID 0)
-- Dependencies: 203
-- Name: product_product_group_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.product_product_group_id_seq OWNED BY public.product.product_group_id;


--
-- TOC entry 205 (class 1259 OID 20147)
-- Name: product_uom_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.product_uom_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.product_uom_seq OWNER TO postgres;

--
-- TOC entry 2725 (class 0 OID 0)
-- Dependencies: 205
-- Name: product_uom_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.product_uom_seq OWNED BY public.product.small_uom_id;


--
-- TOC entry 225 (class 1259 OID 20238)
-- Name: receive; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.receive (
    id bigint NOT NULL,
    receive_no character varying(25),
    receive_date timestamp with time zone,
    supplier_id bigint,
    note character varying(255),
    last_update_by character varying(100),
    last_update timestamp without time zone,
    tax numeric,
    total numeric(19,0),
    grand_total numeric,
    status integer DEFAULT 0
);


ALTER TABLE public.receive OWNER TO postgres;

--
-- TOC entry 224 (class 1259 OID 20236)
-- Name: receive_customer_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.receive_customer_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.receive_customer_id_seq OWNER TO postgres;

--
-- TOC entry 2726 (class 0 OID 0)
-- Dependencies: 224
-- Name: receive_customer_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.receive_customer_id_seq OWNED BY public.receive.supplier_id;


--
-- TOC entry 229 (class 1259 OID 20255)
-- Name: receive_detail; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.receive_detail (
    id bigint NOT NULL,
    receive_id bigint,
    product_id bigint DEFAULT 0 NOT NULL,
    qty numeric DEFAULT 0,
    price numeric DEFAULT 0,
    disc numeric DEFAULT 0,
    uom bigint DEFAULT 0 NOT NULL,
    last_update_by character varying(100),
    last_update timestamp without time zone
);


ALTER TABLE public.receive_detail OWNER TO postgres;

--
-- TOC entry 226 (class 1259 OID 20249)
-- Name: receive_detail_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.receive_detail_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.receive_detail_id_seq OWNER TO postgres;

--
-- TOC entry 2727 (class 0 OID 0)
-- Dependencies: 226
-- Name: receive_detail_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.receive_detail_id_seq OWNED BY public.receive_detail.id;


--
-- TOC entry 227 (class 1259 OID 20251)
-- Name: receive_detail_product_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.receive_detail_product_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.receive_detail_product_id_seq OWNER TO postgres;

--
-- TOC entry 2728 (class 0 OID 0)
-- Dependencies: 227
-- Name: receive_detail_product_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.receive_detail_product_id_seq OWNED BY public.receive_detail.product_id;


--
-- TOC entry 228 (class 1259 OID 20253)
-- Name: receive_detail_uom_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.receive_detail_uom_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.receive_detail_uom_seq OWNER TO postgres;

--
-- TOC entry 2729 (class 0 OID 0)
-- Dependencies: 228
-- Name: receive_detail_uom_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.receive_detail_uom_seq OWNED BY public.receive_detail.uom;


--
-- TOC entry 223 (class 1259 OID 20234)
-- Name: receive_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.receive_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.receive_id_seq OWNER TO postgres;

--
-- TOC entry 2730 (class 0 OID 0)
-- Dependencies: 223
-- Name: receive_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.receive_id_seq OWNED BY public.receive.id;


--
-- TOC entry 233 (class 1259 OID 20312)
-- Name: sales_order_detail; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sales_order_detail (
    id bigint NOT NULL,
    sales_order_id bigint,
    product_id bigint DEFAULT 0 NOT NULL,
    qty real DEFAULT 0,
    price real DEFAULT 0,
    disc real DEFAULT 0,
    uom bigint DEFAULT 0 NOT NULL,
    last_update_by character varying(100),
    last_update timestamp without time zone
);


ALTER TABLE public.sales_order_detail OWNER TO postgres;

--
-- TOC entry 230 (class 1259 OID 20306)
-- Name: sales_detail_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.sales_detail_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.sales_detail_id_seq OWNER TO postgres;

--
-- TOC entry 2731 (class 0 OID 0)
-- Dependencies: 230
-- Name: sales_detail_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.sales_detail_id_seq OWNED BY public.sales_order_detail.id;


--
-- TOC entry 231 (class 1259 OID 20308)
-- Name: sales_detail_product_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.sales_detail_product_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.sales_detail_product_id_seq OWNER TO postgres;

--
-- TOC entry 2732 (class 0 OID 0)
-- Dependencies: 231
-- Name: sales_detail_product_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.sales_detail_product_id_seq OWNED BY public.sales_order_detail.product_id;


--
-- TOC entry 232 (class 1259 OID 20310)
-- Name: sales_detail_uom_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.sales_detail_uom_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.sales_detail_uom_seq OWNER TO postgres;

--
-- TOC entry 2733 (class 0 OID 0)
-- Dependencies: 232
-- Name: sales_detail_uom_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.sales_detail_uom_seq OWNED BY public.sales_order_detail.uom;


--
-- TOC entry 218 (class 1259 OID 20204)
-- Name: supplier; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.supplier (
    id bigint NOT NULL,
    code character varying(10),
    name character varying(100),
    address character varying(255),
    city character varying(100),
    status integer DEFAULT 1,
    last_update_by character varying(100),
    last_update timestamp without time zone,
    tax integer DEFAULT 0,
    pic_name character varying(100),
    pic_phone character varying(50),
    bank_id bigint DEFAULT 0,
    bank_acc_name character varying(100),
    bank_acc_no character varying(50)
);


ALTER TABLE public.supplier OWNER TO postgres;

--
-- TOC entry 217 (class 1259 OID 20202)
-- Name: supplier_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.supplier_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.supplier_id_seq OWNER TO postgres;

--
-- TOC entry 2734 (class 0 OID 0)
-- Dependencies: 217
-- Name: supplier_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.supplier_id_seq OWNED BY public.supplier.id;


--
-- TOC entry 242 (class 1259 OID 20445)
-- Name: tb_sequence_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.tb_sequence_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.tb_sequence_id_seq OWNER TO postgres;

--
-- TOC entry 241 (class 1259 OID 20442)
-- Name: tb_sequence; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.tb_sequence (
    id bigint DEFAULT nextval('public.tb_sequence_id_seq'::regclass),
    subj character varying(10),
    year character varying(4),
    month character varying(2),
    seq integer
);


ALTER TABLE public.tb_sequence OWNER TO postgres;

--
-- TOC entry 2484 (class 2604 OID 20531)
-- Name: adjustment id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.adjustment ALTER COLUMN id SET DEFAULT nextval('public.adjustment_id_seq'::regclass);


--
-- TOC entry 2486 (class 2604 OID 20543)
-- Name: adjustment_detail id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.adjustment_detail ALTER COLUMN id SET DEFAULT nextval('public.adjustment_detail_id_seq'::regclass);


--
-- TOC entry 2446 (class 2604 OID 20163)
-- Name: brand id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.brand ALTER COLUMN id SET DEFAULT nextval('public.brand_id_seq'::regclass);


--
-- TOC entry 2434 (class 2604 OID 20114)
-- Name: customer id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.customer ALTER COLUMN id SET DEFAULT nextval('public.customer_id_seq'::regclass);


--
-- TOC entry 2476 (class 2604 OID 20352)
-- Name: history_stock id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.history_stock ALTER COLUMN id SET DEFAULT nextval('public.history_stock_id_seq'::regclass);


--
-- TOC entry 2439 (class 2604 OID 20135)
-- Name: lookup id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.lookup ALTER COLUMN id SET DEFAULT nextval('public.lookup_id_seq'::regclass);


--
-- TOC entry 2482 (class 2604 OID 20425)
-- Name: lookup_group id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.lookup_group ALTER COLUMN id SET DEFAULT nextval('public.lookup_group_id_seq'::regclass);


--
-- TOC entry 2447 (class 2604 OID 20171)
-- Name: m_menus id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.m_menus ALTER COLUMN id SET DEFAULT nextval('public.m_menus_id_seq'::regclass);


--
-- TOC entry 2448 (class 2604 OID 20179)
-- Name: m_roles id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.m_roles ALTER COLUMN id SET DEFAULT nextval('public.m_roles_id_seq'::regclass);


--
-- TOC entry 2449 (class 2604 OID 20190)
-- Name: m_users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.m_users ALTER COLUMN id SET DEFAULT nextval('public.m_users_id_seq'::regclass);


--
-- TOC entry 2442 (class 2604 OID 20152)
-- Name: product id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.product ALTER COLUMN id SET DEFAULT nextval('public.product_id_seq'::regclass);


--
-- TOC entry 2437 (class 2604 OID 20127)
-- Name: product_group id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.product_group ALTER COLUMN id SET DEFAULT nextval('public.product_group_id_seq'::regclass);


--
-- TOC entry 2462 (class 2604 OID 20241)
-- Name: receive id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.receive ALTER COLUMN id SET DEFAULT nextval('public.receive_id_seq'::regclass);


--
-- TOC entry 2464 (class 2604 OID 20258)
-- Name: receive_detail id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.receive_detail ALTER COLUMN id SET DEFAULT nextval('public.receive_detail_id_seq'::regclass);


--
-- TOC entry 2457 (class 2604 OID 20225)
-- Name: sales_order id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sales_order ALTER COLUMN id SET DEFAULT nextval('public.order_id_seq'::regclass);


--
-- TOC entry 2470 (class 2604 OID 20315)
-- Name: sales_order_detail id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sales_order_detail ALTER COLUMN id SET DEFAULT nextval('public.sales_detail_id_seq'::regclass);


--
-- TOC entry 2453 (class 2604 OID 20207)
-- Name: supplier id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.supplier ALTER COLUMN id SET DEFAULT nextval('public.supplier_id_seq'::regclass);


--
-- TOC entry 2694 (class 0 OID 20528)
-- Dependencies: 244
-- Data for Name: adjustment; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.adjustment (id, adjustment_no, adjustment_date, note, last_update_by, last_update, total, status) FROM stdin;
1	RV19120001	2019-12-15 00:00:00+07	Tes	deddy	2019-12-15 21:42:22.507263	3525000	10
2	AJ191200001	2019-12-15 00:00:00+07		deddy	2019-12-15 22:11:22.028526	0	10
3	AJ191200002	2019-12-15 00:00:00+07		deddy	2019-12-15 22:12:48.124515	4823276	10
\.


--
-- TOC entry 2696 (class 0 OID 20540)
-- Dependencies: 246
-- Data for Name: adjustment_detail; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.adjustment_detail (id, adjustment_id, product_id, qty, hpp, uom, last_update_by, last_update) FROM stdin;
1	1	1	1	2000000	1	system	2019-12-15 00:00:00
2	1	2	-1	1000000	1	system	2019-12-15 00:00:00
3	1	3	1	150000	1		0001-01-01 00:00:00
4	1	1	2	1687500	2		0001-01-01 00:00:00
5	1	2	1	1607758.625	1		0001-01-01 00:00:00
6	1	1	-1	1687500	2		0001-01-01 00:00:00
7	3	2	3	1607758.625	1		0001-01-01 00:00:00
\.


--
-- TOC entry 2658 (class 0 OID 20160)
-- Dependencies: 208
-- Data for Name: brand; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.brand (id, status, code, name, last_update_by, last_update) FROM stdin;
5	1	005	tes lagi		2019-11-14 16:58:01.825315
6	1	006	aaaa		2019-11-14 17:00:21.207735
7	1	007	dsadad		2019-11-14 17:00:27.490674
9	1	009	ddd		2019-11-14 17:00:49.37751
2	1	002	Sharpsssssssss		2019-11-14 17:00:59.182179
11	1	011	teeessssttt	deddy	2019-11-17 17:35:33.373972
1	0	004	LG		2019-11-14 16:57:28.822593
3	0	003	Panasonic		2019-11-14 16:57:28.822593
4	0	001	tesss		2019-11-14 16:58:13.83348
8	1	008	dddd	deddy	2019-11-20 21:06:40.994077
10	1	010	tesssss	deddy	2019-11-21 09:06:50.57485
\.


--
-- TOC entry 2647 (class 0 OID 20111)
-- Dependencies: 197
-- Data for Name: customer; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.customer (id, code, name, last_update_by, last_update, top, status, address1, address2, address3, address4, contact_person, phone_number) FROM stdin;
2	C0000002	Customer 22a	deddy	2019-11-19 09:08:01.040067	30	1	Alamat a	alamat b	alamat c	alamat d	Ibu abcd	08521
3	C000003	testing 	deddy	2019-11-19 09:08:47.507156	3	1	a1	a2	a3	a4	rrr	7765
4	C000004	r	deddy	2019-11-19 09:09:12.862919	22	0	ss	as	ddf	xdff	55	777
1	M000003	Customer 1ss	deddy	2019-11-19 09:09:27.206519	30	0	Alamat 1	alamat 2	alamat 3	alamat 4	Bp 234	999
\.


--
-- TOC entry 2688 (class 0 OID 20349)
-- Dependencies: 238
-- Data for Name: history_stock; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.history_stock (id, code, name, debet, kredit, saldo, trans_date, description, last_update_by, last_update, reff_no, price, hpp) FROM stdin;
1	L00001	LG TV 20 inch	10	0	10	2019-11-10 00:00:00	Purchasing	system	2019-10-11 00:00:00	R191100001	2000000	2000000
2	L00001	LG TV 20 inch	0	1	9	2019-11-10 00:00:00	Penjualan	system	2019-11-10 00:00:00	S191100001	2100000	2000000
3	L00001	LG TV 20 inch	0	1	9	2019-12-01 00:00:00	Sales Order		0001-01-01 00:00:00	SO191200010	1500000	2000000
4	L00001	LG TV 20 inch	0	1	9	2019-12-01 00:00:00	Sales Order		0001-01-01 00:00:00	SO191200010	1500000	2000000
5	L00001	LG TV 20 inch	0	1	8	2019-12-01 00:00:00	Sales Order	deddy	2019-12-01 22:08:30.444572	SO191200010	1500000	2000000
6	L00001	LG TV 20 inch	0	1	8	2019-12-01 00:00:00	Sales Order	deddy	2019-12-01 22:10:33.97792	SO191200010	1500000	2000000
7	L00001	LG TV 20 inch	0	1	8	2019-12-01 00:00:00	Sales Order	deddy	2019-12-01 22:12:20.174953	SO191200010	1500000	2000000
8	L00001	LG TV 20 inch	0	1	8	2019-12-01 00:00:00	Sales Order	deddy	2019-12-01 22:14:34.413062	SO191200010	1500000	2000000
9	L00001	LG TV 20 inch	0	1	8	2019-12-01 00:00:00	Sales Order	deddy	2019-12-01 22:15:04.29453	SO191200010	1500000	2000000
10	L00001	LG TV 20 inch	0	1	8	2019-12-01 00:00:00	Sales Order	deddy	2019-12-01 22:15:52.6363	SO191200010	1500000	2000000
11	L00001	LG TV 20 inch	0	1	8	2019-12-01 00:00:00	Sales Order	deddy	2019-12-01 22:20:40.865632	SO191200010	1500000	2000000
12	L00001	LG TV 20 inch	0	1	7	2019-12-01 00:00:00	Sales Order	deddy	2019-12-01 22:23:54.020611	SO191200010	1500000	2000000
13	L00001	LG TV 20 inch	0	1	6	2019-12-04 00:00:00	Sales Order	deddy	2019-12-04 19:30:40.076329	SO191200007	1500000	2000000
14	L00002	LG Mesin cuci 5kgsaa	0	1	9	2019-12-04 00:00:00	Sales Order	deddy	2019-12-04 19:30:40.086599	SO191200007	1000000	0
15	S00001	Sharp TV 20inch	0	1	9	2019-12-04 00:00:00	Sales Order	deddy	2019-12-04 19:30:40.121445	SO191200007	750000	0
16	P000001	tes ajas	0	3	7	2019-12-04 00:00:00	Sales Order	deddy	2019-12-04 19:30:40.177797	SO191200007	2000000	0
17	L00001	LG TV 20 inch	0	1	5	2019-12-02 00:00:00	Sales Order	deddy	2019-12-04 19:40:09.259259	SO191200011	1500000	2000000
18	L00001	LG TV 20 inch	2	0	7	2019-12-08 00:00:00	Receive	deddy	2019-12-08 20:38:41.193996	RV191200002	1000000	2000000
19	L00002	LG Mesin cuci 5kgsaa	1	0	10	2019-12-08 00:00:00	Receive	deddy	2019-12-08 20:38:41.247778	RV191200002	1500000	0
20	L00001	LG TV 20 inch	1	0	8	2019-12-08 00:00:00	Receive	deddy	2019-12-08 20:44:51.403404	RV191200003	1500000	1714285.75
21	S00001	Sharp TV 20inch	1	0	10	2019-12-08 00:00:00	Receive	deddy	2019-12-08 20:46:26.622686	RV191200001	2500000	0
22	S00001	Sharp TV 20inch	3	0	13	2019-12-15 00:00:00	Receive	deddy	2019-12-15 12:39:20.528419	RV191200005	250000	250000
23	S00001	Sharp TV 20inch	45	0	58	2019-12-15 00:00:00	Receive	deddy	2019-12-15 12:42:52.285162	RV191200004	2000000	250000
24	L00001	LG TV 20 inch	1	0	9	2019-12-15 00:00:00	Adjustment	deddy	2019-12-15 21:42:22.444919	RV19120001	1687500	1687500
25	S00001	Sharp TV 20inch	0	-1	59	2019-12-15 00:00:00	Adjustment	deddy	2019-12-15 21:42:22.478337	RV19120001	1607758.62	1607758.625
26	L00002	LG Mesin cuci 5kgsaa	1	0	11	2019-12-15 00:00:00	Adjustment	deddy	2019-12-15 21:42:22.484655	RV19120001	150000	150000
27	L00001	LG TV 20 inch	2	0	11	2019-12-15 00:00:00	Adjustment	deddy	2019-12-15 21:42:22.490291	RV19120001	1687500	1687500
28	S00001	Sharp TV 20inch	1	0	60	2019-12-15 00:00:00	Adjustment	deddy	2019-12-15 21:42:22.499623	RV19120001	1607758.62	1607758.625
29	L00001	LG TV 20 inch	0	-1	12	2019-12-15 00:00:00	Adjustment	deddy	2019-12-15 21:42:22.504077	RV19120001	1687500	1687500
30	S00001	Sharp TV 20inch	3	0	63	2019-12-15 00:00:00	Adjustment	deddy	2019-12-15 22:12:48.078283	AJ191200002	1607758.62	1607758.625
\.


--
-- TOC entry 2651 (class 0 OID 20132)
-- Dependencies: 201
-- Data for Name: lookup; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.lookup (id, status, code, lookup_group, name, is_viewable) FROM stdin;
4	1	L00004	PAYMENT_METHOD	CASH	1
6	1	L00006	PAYMENT_METHOD	TRANSFER BANK	1
7	1	L00007	PAYMENT_METHOD	GIRO	1
1	1	L00002	SATUAN	liter	1
2	1	L00001	SATUAN	km	1
9	1	L00003	SATUAN	PC	1
5	1	L00003	BANK	BCA	1
3	1	L00003	BANK	Mandiri	1
8	1	L00003	BANK	Permata	1
\.


--
-- TOC entry 2690 (class 0 OID 20422)
-- Dependencies: 240
-- Data for Name: lookup_group; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.lookup_group (id, name) FROM stdin;
1	SATUAN
2	BANK
3	PAYMENT_METHOD
\.


--
-- TOC entry 2660 (class 0 OID 20168)
-- Dependencies: 210
-- Data for Name: m_menus; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.m_menus (id, name, description, last_update_by, last_update, link, parent_id, icon, status) FROM stdin;
7	role	Role	1	2019-09-19 00:00:00	role	22		1
15	user	User	1	2019-09-19 00:00:00	user	22		1
16	access-matrix	Access Matrix	1	2019-09-19 00:00:00	access-matrix	22		1
1	brand	Brand	1	2019-09-19 00:00:00	brand	5		1
2	product-group	Product Group	1	2019-09-19 00:00:00	product-group	5		1
5	master	Master	1	2019-09-19 00:00:00	\N	0		1
11	supplier	Supplier	1	2019-09-19 00:00:00	supplier	5		1
12	product	Product	1	2019-09-19 00:00:00	product	5		1
14	lookup	Lookup	1	2019-09-19 00:00:00	lookup	5		1
3	customer	Customer	1	2019-09-19 00:00:00	customer	5		1
4	dashboard	Dashboard	1	2019-09-19 00:00:00	dashboards	0		1
17	utility	Utility	1	2019-09-19 00:00:00	\N	0		1
18	history	History Stock	1	2019-10-25 00:00:00	history	17		1
20	Adjustment	Adjustment	system	2019-12-15 00:00:00	adjustment	17		1
21	Transaction-Buy	Transaction Buy	system	2019-12-15 00:00:00		0		1
6	transaction	Transaction Sell	1	2019-09-19 00:00:00	\N	0		1
8	sales-order	Sales Order	1	2019-09-19 00:00:00	sales-order	6		1
13	purchase	Purchase Order	1	2019-09-19 00:00:00	purchase	21		1
22	Authorization	Authorization	system	2019-12-15 00:00:00		0		1
\.


--
-- TOC entry 2663 (class 0 OID 20182)
-- Dependencies: 213
-- Data for Name: m_role_menu; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.m_role_menu (role_id, menu_id, status, last_update_by, last_update) FROM stdin;
1	11	1	1	2019-11-01 15:21:54.113
1	19	1	1	2019-11-01 15:21:54.178
1	2	1	1	2019-11-01 15:21:54.242
1	3	1	1	2019-11-01 15:21:54.311
1	5	1	1	2019-11-01 15:21:54.374
1	7	1	1	2019-11-01 15:21:54.437
1	6	1	1	2019-11-01 15:21:54.5
1	4	1	1	2019-11-01 15:21:54.568
1	8	1	1	2019-11-01 15:21:54.631
1	9	1	1	2019-11-01 15:21:54.693
1	17	1	1	2019-11-01 15:21:54.757
1	1	1	1	2019-11-01 15:21:54.823
1	13	1	1	2019-11-01 15:21:54.884
1	12	1	1	2019-11-01 15:21:54.964
1	14	1	1	2019-11-01 15:21:55.026
1	15	1	1	2019-11-01 15:21:55.089
1	18	1	1	2019-11-01 15:21:55.151
1	16	1	1	2019-11-01 15:21:55.213
1	10	1	1	2019-11-01 15:21:55.276
1	20	1	system	2019-12-15 00:00:00
1	21	1	system	2019-12-15 00:00:00
1	22	0	system	2019-12-15 00:00:00
\.


--
-- TOC entry 2666 (class 0 OID 20198)
-- Dependencies: 216
-- Data for Name: m_role_user; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.m_role_user (role_id, user_id, last_update_by, last_update, status) FROM stdin;
1	1	1	2019-09-20 00:00:00	1
\.


--
-- TOC entry 2662 (class 0 OID 20176)
-- Dependencies: 212
-- Data for Name: m_roles; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.m_roles (id, name, description, last_update_by, last_update) FROM stdin;
1	admin	Admin	1	2019-09-16 00:00:00
\.


--
-- TOC entry 2665 (class 0 OID 20187)
-- Dependencies: 215
-- Data for Name: m_users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.m_users (id, user_name, email, password, status, last_update_by, last_update, first_name, last_name, role_id) FROM stdin;
1	deddy	deddy.syuhendra@ottodigital.id	0c91a43f8e1ec5fcba28f8a5a34532679305ca131302ad2a06218b47f30ced88	1	deddy	2019-09-16 00:00:00	deddy	syuhendra	1
\.


--
-- TOC entry 2656 (class 0 OID 20149)
-- Dependencies: 206
-- Data for Name: product; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.product (id, code, name, product_group_id, brand_id, small_uom_id, status, last_update_by, last_update, big_uom_id, qty_uom, qty_stock, hpp, sell_price) FROM stdin;
4	P000001	tes ajas	4	3	9	1	deddy	2019-11-21 09:04:14.782048	9	1	7	0	2000000
3	L00002	LG Mesin cuci 5kgsaa	5	1	1	1		2019-11-20 22:13:58.366193	2	2	11	150000	1000000
1	L00001	LG TV 20 inch	2	1	2	1		2019-11-20 22:42:54.389353	1	1	12	1687500	1500000
2	S00001	Sharp TV 20inch	4	3	1	1	deddy	2019-11-21 09:07:10.507013	9	1	63	1607758.62	750000
\.


--
-- TOC entry 2649 (class 0 OID 20124)
-- Dependencies: 199
-- Data for Name: product_group; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.product_group (id, name, status, last_update_by, last_update, code) FROM stdin;
3	AC Indoorsaa	1	deddy	2019-11-17 17:36:18.124115	G003
4	TV	1	deddy	2019-11-17 17:36:18.124115	G001
2	AC Stand	1	deddy	2019-11-17 17:36:18.124115	G002
1	Dish washer	1	deddy	2019-11-17 17:36:18.124115	G004
5	shirt washer	1	deddy	2019-11-17 17:36:18.124115	G005
\.


--
-- TOC entry 2675 (class 0 OID 20238)
-- Dependencies: 225
-- Data for Name: receive; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.receive (id, receive_no, receive_date, supplier_id, note, last_update_by, last_update, tax, total, grand_total, status) FROM stdin;
1		2019-12-08 00:00:00+07	0	Pemesanan tv lg 10pcs	deddy	2019-12-08 20:00:20.207134	2000000	20000000	22000000	10
3	RV191200002	2019-12-08 00:00:00+07	2		deddy	2019-12-08 20:38:41.253835	0	3500000	0	20
4	RV191200003	2019-12-08 00:00:00+07	2		deddy	2019-12-08 20:44:51.415981	0	1500000	0	20
2	RV191200001	2019-12-08 00:00:00+07	2		deddy	2019-12-08 20:46:26.64384	0	2500000	0	20
6	RV191200005	2019-12-15 00:00:00+07	1		deddy	2019-12-15 12:39:20.575596	0	750000	0	20
5	RV191200004	2019-12-15 00:00:00+07	2		deddy	2019-12-15 12:42:52.293984	0	90000000	0	20
\.


--
-- TOC entry 2679 (class 0 OID 20255)
-- Dependencies: 229
-- Data for Name: receive_detail; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.receive_detail (id, receive_id, product_id, qty, price, disc, uom, last_update_by, last_update) FROM stdin;
2	1	3	1	10000	0	1		0001-01-01 00:00:00
3	3	1	2	1000000	0	2		0001-01-01 00:00:00
4	3	3	1	1500000	0	1		0001-01-01 00:00:00
5	4	1	1	1500000	0	2		0001-01-01 00:00:00
6	2	2	1	2500000	0	1		0001-01-01 00:00:00
7	5	2	45	2000000	0	1		0001-01-01 00:00:00
8	6	2	3	250000	0	1		0001-01-01 00:00:00
\.


--
-- TOC entry 2672 (class 0 OID 20222)
-- Dependencies: 222
-- Data for Name: sales_order; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sales_order (id, sales_order_no, order_date, customer_id, note, last_update_by, last_update, tax, total, grand_total, salesman_id, status, top, is_cash) FROM stdin;
2	S191100002	2019-11-20 00:00:00+07	1			2019-11-21 00:00:00	1000	10000	11000	1	1	0	1
3	S191100003	2019-11-21 00:00:00+07	1			2019-11-21 00:00:00	2000	20000	22000	1	1	0	1
4	S191100004	2019-11-22 00:00:00+07	1			2019-11-21 00:00:00	3000	30000	33000	1	1	10	0
5	SO191200003	0001-01-01 07:07:12+07:07:12	1		deddy	2019-12-01 12:52:52.164328	0	0	0	0	10	0	0
1	S191100001	2019-12-01 00:00:00+07	1	Dikirim tgl 10	deddy	2019-12-01 13:02:24.950723	210000	2100000	2310000	0	1	1	1
6	SO191200004	2019-12-01 00:00:00+07	1		deddy	2019-12-01 13:02:47.213501	0	0	0	0	10	0	0
7	SO191200005	2019-12-01 00:00:00+07	1		deddy	2019-12-01 13:03:20.004803	0	0	0	0	10	0	0
8	SO191200006	2019-12-01 00:00:00+07	1		deddy	2019-12-01 13:03:27.947432	0	0	0	0	10	0	0
10	SO191200008	2020-01-01 00:00:00+07	2		deddy	2019-12-01 13:09:16.400295	0	0	0	0	10	0	0
12	SO191200010	2019-12-01 00:00:00+07	1		deddy	2019-12-01 22:23:54.036735	0	1500000	0	1	20	0	0
11	SO191200009	2019-12-05 00:00:00+07	2			2019-12-01 20:01:53.397537	0	0	0	0	30	0	0
14	SO191200012	2019-12-02 00:00:00+07	1		deddy	2019-12-02 20:14:28.009998	0	0	0	0	10	0	0
9	SO191200007	2019-12-04 00:00:00+07	1		deddy	2019-12-04 19:30:40.213119	0	9250000	0	1	20	0	0
13	SO191200011	2019-12-02 00:00:00+07	2		deddy	2019-12-04 19:40:09.264464	0	1500000	0	1	20	0	0
\.


--
-- TOC entry 2683 (class 0 OID 20312)
-- Dependencies: 233
-- Data for Name: sales_order_detail; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sales_order_detail (id, sales_order_id, product_id, qty, price, disc, uom, last_update_by, last_update) FROM stdin;
1	1	1	1	2100000	0	1	system	2019-11-10 00:00:00
4	1	1	1	1500000	1000	2		0001-01-01 00:00:00
5	1	2	2	750000	0	1		0001-01-01 00:00:00
7	1	2	1	750000	0	1		0001-01-01 00:00:00
8	1	1	1	1500000	0	2		0001-01-01 00:00:00
10	1	4	1	2000000	0	9		0001-01-01 00:00:00
11	1	2	1	750000	0	1		0001-01-01 00:00:00
18	12	1	1	1500000	0	2		0001-01-01 00:00:00
19	11	1	1	1500000	0	2		0001-01-01 00:00:00
20	11	2	10	750000	0	1		0001-01-01 00:00:00
21	14	2	10	750000	0	1		0001-01-01 00:00:00
22	10	3	7	1000000	0	1		0001-01-01 00:00:00
23	10	2	6	750000	0	1		0001-01-01 00:00:00
24	9	1	1	1500000	0	2		0001-01-01 00:00:00
25	9	3	1	1000000	0	1		0001-01-01 00:00:00
26	9	2	1	750000	0	1		0001-01-01 00:00:00
27	9	4	3	2000000	0	9		0001-01-01 00:00:00
28	13	1	1	1500000	0	2		0001-01-01 00:00:00
\.


--
-- TOC entry 2668 (class 0 OID 20204)
-- Dependencies: 218
-- Data for Name: supplier; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.supplier (id, code, name, address, city, status, last_update_by, last_update, tax, pic_name, pic_phone, bank_id, bank_acc_name, bank_acc_no) FROM stdin;
2	S002	Sharp Indonesiass	Jl sha	Palembang	1	deddy	2019-11-19 11:52:49.374653	1	dono	0811111	3	tes 2	443322
5	S001	teestsh	hjkhj	hkj	1		2019-11-19 11:58:06.061232	78	hjhj	hjhjjh	8	jh	hjhjjh
6	S001	jhkhh	kjhjh	kjhkhhj	1		2019-11-19 11:58:25.990001	887	ljkjkjkl	uiljkjk	8	kjjkl	ljkkjl
7	S001	dadsa	ddada	dasd	1		2019-11-19 11:59:56.826319	213	asda	dasd	8	dasda	dasd
8	S001	luiuou	iooiu	uuiooi	1		2019-11-19 12:02:19.343339	87	ljkjkl	jkjklj	8	jkl	jjkljkl
9	002	jkhkjk	jhhjjkh	jkhjjk	1		2019-11-19 12:04:24.006144	88	jkh	jhhjjhk	8	jhk	hjkhkj
10	S003	uuuhuiii	uuu	uuu	1		2019-11-19 12:05:55.705282	88	ioioio	iiioio	8	iopoi	iopiopio
11	S004	dadsada	dada	dasdas	1		2019-11-19 12:09:56.191859	233	asdasda	dsa	5	dasda	dasda
1	S001	LG Indonesia	Jl Elgi	Palembang	1		2019-11-20 22:51:59.51184	1	jojon	08123456	8	tes 1	112233
3	S003	Polytron Indonesia	Jl Pol	Palembang	1		2019-11-20 22:52:22.954466	1	sasa	085666	3	tes 3	4213131
\.


--
-- TOC entry 2691 (class 0 OID 20442)
-- Dependencies: 241
-- Data for Name: tb_sequence; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.tb_sequence (id, subj, year, month, seq) FROM stdin;
1	SO	2019	11	1
2	SO	2019	12	4
3	SO	19	12	12
4	RV	19	12	5
5	AJ	19	12	2
\.


--
-- TOC entry 2735 (class 0 OID 0)
-- Dependencies: 245
-- Name: adjustment_detail_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.adjustment_detail_id_seq', 7, true);


--
-- TOC entry 2736 (class 0 OID 0)
-- Dependencies: 243
-- Name: adjustment_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.adjustment_id_seq', 3, true);


--
-- TOC entry 2737 (class 0 OID 0)
-- Dependencies: 207
-- Name: brand_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.brand_id_seq', 11, true);


--
-- TOC entry 2738 (class 0 OID 0)
-- Dependencies: 196
-- Name: customer_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.customer_id_seq', 4, true);


--
-- TOC entry 2739 (class 0 OID 0)
-- Dependencies: 235
-- Name: history_stock_debet_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.history_stock_debet_seq', 1, false);


--
-- TOC entry 2740 (class 0 OID 0)
-- Dependencies: 234
-- Name: history_stock_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.history_stock_id_seq', 30, true);


--
-- TOC entry 2741 (class 0 OID 0)
-- Dependencies: 236
-- Name: history_stock_kredit_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.history_stock_kredit_seq', 1, false);


--
-- TOC entry 2742 (class 0 OID 0)
-- Dependencies: 237
-- Name: history_stock_saldo_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.history_stock_saldo_seq', 1, false);


--
-- TOC entry 2743 (class 0 OID 0)
-- Dependencies: 239
-- Name: lookup_group_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.lookup_group_id_seq', 3, true);


--
-- TOC entry 2744 (class 0 OID 0)
-- Dependencies: 200
-- Name: lookup_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.lookup_id_seq', 9, true);


--
-- TOC entry 2745 (class 0 OID 0)
-- Dependencies: 209
-- Name: m_menus_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.m_menus_id_seq', 22, true);


--
-- TOC entry 2746 (class 0 OID 0)
-- Dependencies: 211
-- Name: m_roles_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.m_roles_id_seq', 1, true);


--
-- TOC entry 2747 (class 0 OID 0)
-- Dependencies: 214
-- Name: m_users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.m_users_id_seq', 1, true);


--
-- TOC entry 2748 (class 0 OID 0)
-- Dependencies: 220
-- Name: order_customer_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.order_customer_id_seq', 1, false);


--
-- TOC entry 2749 (class 0 OID 0)
-- Dependencies: 219
-- Name: order_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.order_id_seq', 14, true);


--
-- TOC entry 2750 (class 0 OID 0)
-- Dependencies: 221
-- Name: order_salesman_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.order_salesman_id_seq', 1, false);


--
-- TOC entry 2751 (class 0 OID 0)
-- Dependencies: 204
-- Name: product_brand_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.product_brand_id_seq', 3, true);


--
-- TOC entry 2752 (class 0 OID 0)
-- Dependencies: 198
-- Name: product_group_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.product_group_id_seq', 5, true);


--
-- TOC entry 2753 (class 0 OID 0)
-- Dependencies: 202
-- Name: product_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.product_id_seq', 4, true);


--
-- TOC entry 2754 (class 0 OID 0)
-- Dependencies: 203
-- Name: product_product_group_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.product_product_group_id_seq', 3, true);


--
-- TOC entry 2755 (class 0 OID 0)
-- Dependencies: 205
-- Name: product_uom_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.product_uom_seq', 3, true);


--
-- TOC entry 2756 (class 0 OID 0)
-- Dependencies: 224
-- Name: receive_customer_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.receive_customer_id_seq', 1, false);


--
-- TOC entry 2757 (class 0 OID 0)
-- Dependencies: 226
-- Name: receive_detail_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.receive_detail_id_seq', 8, true);


--
-- TOC entry 2758 (class 0 OID 0)
-- Dependencies: 227
-- Name: receive_detail_product_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.receive_detail_product_id_seq', 1, false);


--
-- TOC entry 2759 (class 0 OID 0)
-- Dependencies: 228
-- Name: receive_detail_uom_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.receive_detail_uom_seq', 1, false);


--
-- TOC entry 2760 (class 0 OID 0)
-- Dependencies: 223
-- Name: receive_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.receive_id_seq', 6, true);


--
-- TOC entry 2761 (class 0 OID 0)
-- Dependencies: 230
-- Name: sales_detail_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.sales_detail_id_seq', 28, true);


--
-- TOC entry 2762 (class 0 OID 0)
-- Dependencies: 231
-- Name: sales_detail_product_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.sales_detail_product_id_seq', 1, false);


--
-- TOC entry 2763 (class 0 OID 0)
-- Dependencies: 232
-- Name: sales_detail_uom_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.sales_detail_uom_seq', 1, false);


--
-- TOC entry 2764 (class 0 OID 0)
-- Dependencies: 217
-- Name: supplier_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.supplier_id_seq', 11, true);


--
-- TOC entry 2765 (class 0 OID 0)
-- Dependencies: 242
-- Name: tb_sequence_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.tb_sequence_id_seq', 5, true);


--
-- TOC entry 2524 (class 2606 OID 20552)
-- Name: adjustment_detail adjustment_detail_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.adjustment_detail
    ADD CONSTRAINT adjustment_detail_pkey PRIMARY KEY (id);


--
-- TOC entry 2522 (class 2606 OID 20537)
-- Name: adjustment adjustment_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.adjustment
    ADD CONSTRAINT adjustment_pkey PRIMARY KEY (id);


--
-- TOC entry 2500 (class 2606 OID 20165)
-- Name: brand brand_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.brand
    ADD CONSTRAINT brand_pkey PRIMARY KEY (id);


--
-- TOC entry 2518 (class 2606 OID 20360)
-- Name: history_stock history_stock_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.history_stock
    ADD CONSTRAINT history_stock_pkey PRIMARY KEY (id);


--
-- TOC entry 2520 (class 2606 OID 20427)
-- Name: lookup_group lookup_group_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.lookup_group
    ADD CONSTRAINT lookup_group_pkey PRIMARY KEY (id);


--
-- TOC entry 2496 (class 2606 OID 20140)
-- Name: lookup lookup_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.lookup
    ADD CONSTRAINT lookup_pkey PRIMARY KEY (id);


--
-- TOC entry 2502 (class 2606 OID 20173)
-- Name: m_menus m_menus_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.m_menus
    ADD CONSTRAINT m_menus_pkey PRIMARY KEY (id);


--
-- TOC entry 2504 (class 2606 OID 20181)
-- Name: m_roles m_roles_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.m_roles
    ADD CONSTRAINT m_roles_pkey PRIMARY KEY (id);


--
-- TOC entry 2506 (class 2606 OID 20197)
-- Name: m_users m_users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.m_users
    ADD CONSTRAINT m_users_pkey PRIMARY KEY (id);


--
-- TOC entry 2492 (class 2606 OID 20121)
-- Name: customer merchant_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.customer
    ADD CONSTRAINT merchant_pkey PRIMARY KEY (id);


--
-- TOC entry 2510 (class 2606 OID 20233)
-- Name: sales_order order_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sales_order
    ADD CONSTRAINT order_pkey PRIMARY KEY (id);


--
-- TOC entry 2494 (class 2606 OID 20129)
-- Name: product_group product_group_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.product_group
    ADD CONSTRAINT product_group_pkey PRIMARY KEY (id);


--
-- TOC entry 2498 (class 2606 OID 20157)
-- Name: product product_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.product
    ADD CONSTRAINT product_pkey PRIMARY KEY (id);


--
-- TOC entry 2514 (class 2606 OID 20268)
-- Name: receive_detail receive_detail_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.receive_detail
    ADD CONSTRAINT receive_detail_pkey PRIMARY KEY (id);


--
-- TOC entry 2512 (class 2606 OID 20248)
-- Name: receive receive_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.receive
    ADD CONSTRAINT receive_pkey PRIMARY KEY (id);


--
-- TOC entry 2516 (class 2606 OID 20325)
-- Name: sales_order_detail sales_detail_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sales_order_detail
    ADD CONSTRAINT sales_detail_pkey PRIMARY KEY (id);


--
-- TOC entry 2508 (class 2606 OID 20214)
-- Name: supplier supplier_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.supplier
    ADD CONSTRAINT supplier_pkey PRIMARY KEY (id);


-- Completed on 2019-12-18 20:16:21 WIB

--
-- PostgreSQL database dump complete
--

