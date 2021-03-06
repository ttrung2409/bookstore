package com.bookstore.ecommerce.app.repository;

import java.util.concurrent.CompletableFuture;

public interface TransactionFactory {
  <R> CompletableFuture<R> runInTransaction(TransactionalFunc<CompletableFuture<R>> func)
    throws Exception;
}
