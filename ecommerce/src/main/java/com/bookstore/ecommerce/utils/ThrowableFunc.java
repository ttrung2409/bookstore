package com.bookstore.ecommerce.utils;

@FunctionalInterface
public interface ThrowableFunc<R> {
  R apply() throws Exception;
}

