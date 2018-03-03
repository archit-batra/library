//package library;
import java.util.*;
import java.sql.PreparedStatement;
import java.sql.DriverManager;
import java.sql.Connection;
import java.sql.ResultSet;
public class library {
    public static void main(String[] args) {
        boolean loop = true;
        do{
        String mainMenu = ("WELCOME TO LIBRARY. MAINTAIN SILENCE\n"+
        "WHAT DO YOU WANT TO DO?\n"+
        "1 --> LIST BOOKS\n"+
        "2 --> SEARCH BOOKS\n"+
        "3 --> ADD BOOKS\n"+
        "4 --> DELETE BOOKS\n"+
        "5 --> EXIT");
        System.out.println(mainMenu);
        Connection con;
        PreparedStatement stmt;
        ResultSet rs;
        String sql;
        String BookName;
        String AuthorName;
        int BookId;
        Scanner sc = new Scanner(System.in);
        int choice = sc.nextInt();
        switch(choice)
{
case 1:
        try{        
            Class.forName("com.mysql.jdbc.Driver"); 
            con = DriverManager.getConnection("jdbc:mysql://localhost:3306/library","root","");
            sql = "select * from books";
            stmt = con.prepareStatement(sql);
            rs=stmt.executeQuery();
            while(rs.next())
            System.out.println(rs.getInt(1)+" "+rs.getString(2)+" by "+rs.getString(3));   
            con.close();  
	   }
        catch(Exception e){
            e.getStackTrace();
                          }
break;
case 2:
        System.out.println("Enter the name of book and author you want to search:");
        sc.nextLine();
        BookName = sc.nextLine();
        AuthorName = sc.nextLine();
        try{        
            Class.forName("com.mysql.jdbc.Driver");
            con = DriverManager.getConnection("jdbc:mysql://localhost:3306/library","root","");
            sql = "select * from books where name = ? and author = ?";
            stmt = con.prepareStatement(sql);
            stmt.setString(1, BookName);
            stmt.setString(2, AuthorName);
            rs=stmt.executeQuery();
            if(rs.next())
            System.out.println("book found");
            else
            System.out.println("book not found");
            con.close();
            } 
        catch(Exception e){
            System.out.println(e.getMessage());                                   
                          } 
break;
case 3:
       System.out.println("Please enter book id, name of book and author you want to add:");
       BookId = sc.nextInt();
       sc.nextLine();
       BookName = sc.nextLine();
       AuthorName = sc.nextLine();
       try{        
            Class.forName("com.mysql.jdbc.Driver");
            con = DriverManager.getConnection("jdbc:mysql://localhost:3306/library","root","");
            sql="insert into books(id,name,author) values(?,?,?)";
            stmt = con.prepareStatement(sql);
            stmt.setInt(1, BookId);
            stmt.setString(2, BookName);
            stmt.setString(3, AuthorName);
            int b=stmt.executeUpdate();
            if(b!=0)
            	System.out.println(BookId+" "+BookName+" by "+AuthorName+" added to database");
            else
                System.out.println("book not added");
            con.close();
            } 
       catch (Exception e){
       System.out.println(e.getMessage());
                          } 
break;
case 4:
       System.out.println("Please enter book id you want to delete:");
       BookId = sc.nextInt();
       try{        
            Class.forName("com.mysql.jdbc.Driver");
            con = DriverManager.getConnection("jdbc:mysql://localhost:3306/library","root","");
            sql="delete from books where id = ?";
            stmt = con.prepareStatement(sql);
            stmt.setInt(1, BookId);
            int b=stmt.executeUpdate();
            if(b!=0)
            	System.out.println("Book id "+BookId+" deleted from database");
            else
                System.out.println("book not deleted");
            con.close();
            } 
       catch (Exception e){
       System.out.println(e.getMessage());
                          } 
break;
case 5:
    System.exit(0);
default:
    System.out.println("\nError! Incorrect choice.\n");
    break;
        }
        }while(loop==true);
        }
    }
                    