package com.bookstore.ecommerce.app.repository.query;

import java.util.List;
import com.bookstore.ecommerce.app.domain.data.Book;
import com.bookstore.ecommerce.app.repository.Query;

import lombok.Builder;
import lombok.Data;

public interface BooksQuery extends Query<BooksQuery.Params, List<Book>> {

  @Data
  @Builder
  public class Params {
    String term;
  }
}
