package com.bookstore.ecommerce.app.order.query;

import java.util.concurrent.CompletableFuture;
import com.bookstore.ecommerce.app.order.query.dto.Book;
import com.bookstore.ecommerce.app.order.query.dto.Order;

public interface Query {
  CompletableFuture<Book[]> findBooks(String term);

  Order getOrderDetails(String orderId);
}
