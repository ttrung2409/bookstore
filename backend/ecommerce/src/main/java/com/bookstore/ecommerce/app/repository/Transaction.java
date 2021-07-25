package com.bookstore.ecommerce.app.repository;

public interface Transaction {
  void commit() throws Exception;

  void rollback() throws Exception;
}
