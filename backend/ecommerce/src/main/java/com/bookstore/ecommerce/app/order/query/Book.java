package com.bookstore.ecommerce.app.order.query;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;

@AllArgsConstructor
@Builder
@Data
public class Book {
  private String id;
  private String title;
  private String subTitle;
  private String description;
}
