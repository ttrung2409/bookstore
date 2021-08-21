package com.bookstore.ecommerce.app.operation.query;

import java.util.concurrent.CompletableFuture;

public interface BookQuery {
  CompletableFuture<Book[]> find(String term);
}
