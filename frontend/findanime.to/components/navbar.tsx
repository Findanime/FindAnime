import React from "react";
import {
  Navbar as HeroUINavbar,
  NavbarContent,
  NavbarBrand,
  NavbarItem,
  NavbarMenu,
  NavbarMenuToggle,
} from "@heroui/navbar";
import { Link } from "@heroui/link";
import NextLink from "next/link";
import clsx from "clsx";

export const Navbar = () => {
  return (
    <div className="w-full flex justify-center">
      <HeroUINavbar
        maxWidth="xl"
        position="sticky"
        className="w-full max-w-7xl px-6"
      >
        <NavbarContent className="basis-1/5 sm:basis-full" justify="start">
          <NavbarBrand as="li" className="gap-3 max-w-fit">
            <NextLink
              className="flex justify-start items-center gap-1"
              href="/"
            >
              <p className="font-bold text-inherit">FindAnime</p>
            </NextLink>
          </NavbarBrand>

          <ul className="hidden lg:flex gap-4 justify-start ml-2">
            <NavbarItem>
              <NextLink
                className={clsx(
                  "text-white",
                  "hover:text-primary",
                  "font-medium"
                )}
                href="/"
              >
                Home
              </NextLink>
            </NavbarItem>
            <NavbarItem>
              <NextLink
                className={clsx(
                  "text-white",
                  "hover:text-primary",
                  "font-medium"
                )}
                href="/movies"
              >
                Movies
              </NextLink>
            </NavbarItem>
            <NavbarItem>
              <NextLink
                className={clsx(
                  "text-white",
                  "hover:text-primary",
                  "font-medium"
                )}
                href="/tv-series"
              >
                TV Series
              </NextLink>
            </NavbarItem>
            <NavbarItem>
              <NextLink
                className={clsx(
                  "text-white",
                  "hover:text-primary",
                  "font-medium"
                )}
                href="/most-popular"
              >
                Most Popular
              </NextLink>
            </NavbarItem>
            <NavbarItem>
              <NextLink
                className={clsx(
                  "text-white",
                  "hover:text-primary",
                  "font-medium"
                )}
                href="/top-airing"
              >
                Top Airing
              </NextLink>
            </NavbarItem>
          </ul>
        </NavbarContent>

        <NavbarContent
          className="hidden sm:flex basis-1/5 sm:basis-full"
          justify="end"
        >
          <NavbarItem className="hidden sm:flex gap-2">
            {/* Buttons */}
          </NavbarItem>
        </NavbarContent>

        <NavbarContent className="sm:hidden basis-1 pl-4" justify="end">
          <NavbarMenuToggle />
        </NavbarContent>

        <NavbarMenu>{/* Mobile menu */}</NavbarMenu>
      </HeroUINavbar>
    </div>
  );
};
