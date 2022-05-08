package com.bookstore.ecommerce.app.repository;

import java.util.concurrent.CompletableFuture;
import com.bookstore.ecommerce.app.domain.Customer;

public interface CustomerRepository {
  CompletableFuture<Void> createIfNotExist(Customer customer, Transaction tx) throws Exception;
}
