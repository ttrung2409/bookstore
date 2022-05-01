package com.bookstore.ecommerce.utils;

@FunctionalInterface
public interface ThrowableSupplier<T> {
  T get() throws Exception;
}

