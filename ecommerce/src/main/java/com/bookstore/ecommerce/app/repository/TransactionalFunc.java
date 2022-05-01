package com.bookstore.ecommerce.app.repository;

@FunctionalInterface
public interface TransactionalFunc<R> {
  R apply(Transaction tx) throws Exception;
}
