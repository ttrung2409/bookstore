package com.bookstore.ecommerce.app.repository;

import java.util.concurrent.CompletableFuture;

public interface Query<TParams, TResult> {
  CompletableFuture<TResult> execute(TParams params);
}
