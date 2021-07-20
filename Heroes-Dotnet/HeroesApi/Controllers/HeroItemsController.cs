using System;
using System.Collections.Generic;
using System.Linq;
using HeroesApi.Models;
using Microsoft.AspNetCore.Authorization;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;

namespace HeroesApi.Controllers
{
    [ApiController]
    [Authorize]
    [Route("api/heroitems")]
    public class HeroItemsController : ControllerBase
    {
        private readonly HeroContext _context;

        public HeroItemsController(HeroContext context)
        {
            _context = context;
        }

        // GET: api/heroitems
        [HttpGet]
        public async Task<ActionResult<IEnumerable<HeroItem>>> GetHeroItems()
        {
            return await _context.HeroItems.ToListAsync();
        }

        // GET: api/heroitems/5
        [HttpGet("{id}")]
        public async Task<ActionResult<HeroItem>> GetHeroItem(long id)
        {
            var heroItem = await _context.HeroItems.FindAsync(id);

            if (heroItem == null)
            {
                return NotFound();
            }

            return heroItem;
        }

        // PUT: api/heroitems/5
        [HttpPut("{id}")]
        public async Task<IActionResult> PutHeroItem(long id, HeroItem heroItem)
        {
            if (id != heroItem.Id)
            {
                return BadRequest();
            }

            _context.Entry(heroItem).State = EntityState.Modified;

            try
            {
                await _context.SaveChangesAsync();
            }
            catch (DbUpdateConcurrencyException)
            {
                if (!HeroItemExists(id))
                {
                    return NotFound();
                }
                else
                {
                    throw;
                }
            }

            return NoContent();
        }

        // POST: api/heroitems
        [HttpPost]
        public async Task<ActionResult<HeroItem>> PostHeroItem(HeroItem heroItem)
        {
            _context.HeroItems.Add(heroItem);
            await _context.SaveChangesAsync();

            return CreatedAtAction("GetHeroItem", new { id = heroItem.Id }, heroItem);
        }

        // DELETE: api/heroitems/5
        [HttpDelete("{id}")]
        public async Task<ActionResult<HeroItem>> DeleteHeroItem(long id)
        {
            var heroItem = await _context.HeroItems.FindAsync(id);
            if (heroItem == null)
            {
                return NotFound();
            }

            _context.HeroItems.Remove(heroItem);
            await _context.SaveChangesAsync();

            return heroItem;
        }

        private bool HeroItemExists(long id)
        {
            return _context.HeroItems.Any(e => e.Id == id);
        }
    }
}
