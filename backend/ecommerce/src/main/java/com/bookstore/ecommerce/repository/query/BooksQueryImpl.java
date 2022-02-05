package com.bookstore.ecommerce.repository.query;

import java.util.List;
import java.util.concurrent.CompletableFuture;
import com.bookstore.ecommerce.app.domain.data.Book;
import com.bookstore.ecommerce.app.repository.query.BooksQuery;

public class BooksQueryImpl implements BooksQuery {

  @Override
  public CompletableFuture<List<Book>> execute(Params params) {
    throw new Error("Not implemented");
  }

}
