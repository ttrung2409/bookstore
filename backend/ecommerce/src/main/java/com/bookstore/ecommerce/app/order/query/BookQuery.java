package com.bookstore.ecommerce.app.order.query;

import java.util.concurrent.CompletableFuture;

public interface BookQuery {
  CompletableFuture<Book[]> find(String term);
}
