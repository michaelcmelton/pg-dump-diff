-- Dumped from database version 12.3
-- Dumped by pg_dump version 12.3

SET statement_timeout = '0';
SET lock_timeout = '0';
SET idle_in_transaction_session_timeout = '0';
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;

-- Schema for public
CREATE SCHEMA public;

-- Table public.users
CREATE TABLE public.users (
                              id SERIAL PRIMARY KEY,
                              username VARCHAR(50) NOT NULL UNIQUE,
                              password VARCHAR(255) NOT NULL,
                              created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Table public.posts
CREATE TABLE public.posts (
                              id SERIAL PRIMARY KEY,
                              user_id INTEGER NOT NULL REFERENCES public.users(id) ON DELETE CASCADE,
                              title VARCHAR(100) NOT NULL,
                              content TEXT NOT NULL,
                              created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Indexes
CREATE INDEX idx_users_username ON public.users (username);
CREATE INDEX idx_posts_user_id ON public.posts (user_id);

-- Constraints
ALTER TABLE public.posts ADD CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES public.users(id);
