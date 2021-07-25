package com.bookstore.ecommerce.app.repository;

import java.util.concurrent.Callable;
import java.util.concurrent.CompletableFuture;

public interface TransactionFactory {
  <T> T runInTransaction(Callable<CompletableFuture<T>> func) throws Exception;
}
